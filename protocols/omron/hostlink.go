package omron

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/protocol/helper"
	"time"
)

type FinsHostLink struct {
	frame UdpFrame
	link  connect.Link
	queue chan *request //in
}

func NewFinsHostLink(link connect.Link) *FinsHostLink {
	fins := &FinsHostLink{
		link:  link,
		queue: make(chan *request, 1),
	}
	link.On("data", func(data []byte) {
		fins.OnData(data)
	})
	link.On("close", func() {
		close(fins.queue)
	})
	return fins
}

func (f *FinsHostLink) execute(cmd []byte) ([]byte, error) {
	req := &request{
		cmd:  cmd,
		resp: make(chan response, 1),
	}
	//排队等待
	f.queue <- req

	//下发指令
	err := f.link.Write(cmd)
	if err != nil {
		//释放队列
		<-f.queue
		return nil, err
	}

	//等待结果
	select {
	case <-time.After(5 * time.Second):
		<-f.queue //清空
		return nil, errors.New("timeout")
	case resp := <-req.resp:
		return resp.buf, resp.err
	}
}

func (f *FinsHostLink) OnData(buf []byte) {
	if len(f.queue) == 0 {
		//无效数据
		return
	}

	//取出请求，并让出队列，可以开始下一个请示了
	req := <-f.queue

	//解析数据
	l := len(buf)
	if l < 23 {
		return
	}

	//@ [单元号] [F A] [0 0] [4 0 ICF][0 0 DA2][0 0 SA2][ SID ]
	//[命令码 4字节] [状态码 4字节] [ ...data... ]
	//[FCS][* CR]
	recv := helper.FromHex(buf[15 : l-4])

	//记录响应的SID
	//t.frame.SID = FromHex(payload[13:15])[0]

	req.resp <- response{buf: recv}
}

func (f *FinsHostLink) Address(addr string) (protocol.Addr, error) {
	return ParseAddress(addr)
}

func (f *FinsHostLink) Read(station int, address protocol.Addr, size int) ([]byte, error) {

	//构建读命令
	buf, e := buildReadCommand(address, size)
	if e != nil {
		return nil, e
	}

	//打包命令
	cmd := packAsciiCommand(&f.frame, buf)

	//发送请求
	recv, err := f.execute(cmd)
	if err != nil {
		return nil, err
	}

	//[命令码 1 1] [结束码 0 0] , data
	code := helper.ParseUint16(recv[2:])
	if code != 0 {
		return nil, fmt.Errorf("错误码: %d", code)
	}

	return recv[4:], nil
}

func (f *FinsHostLink) Immediate(station int, addr protocol.Addr, size int) ([]byte, error) {
	return f.Read(station, addr, size)
}

func (f *FinsHostLink) Write(station int, address protocol.Addr, values []byte) error {
	//构建写命令
	buf, e := buildWriteCommand(address, values)
	if e != nil {
		return e
	}

	//打包命令
	cmd := packAsciiCommand(&f.frame, buf)

	//发送请求
	recv, err := f.execute(cmd)
	if err != nil {
		return err
	}

	//[命令码 1 1] [结束码 0 0]
	code := helper.ParseUint16(recv[2:])
	if code != 0 {
		return fmt.Errorf("错误码: %d", code)
	}

	return nil
}
