package omron

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/protocol/helper"
	"github.com/zgwit/iot-master/connect"
)

type FinsHostLink struct {
	frame UdpFrame
	link  connect.Link
}

func NewFinsHostLink(link connect.Link) *FinsHostLink {
	a := &FinsHostLink{}
	a.link = link
	return a
}

func (t *FinsHostLink) request(cmd []byte) ([]byte, error) {
	if e := t.link.Write(cmd); e != nil {
		return nil, e
	}

	payload := make([]byte, 1024)
	n := 80 //fake
	//n, err := t.link.Read(payload)
	//if err != nil {
	//	return nil, err
	//}

	//@ [单元号] [F A] [0 0] [4 0 ICF][0 0 DA2][0 0 SA2][ SID ]
	//[命令码 4字节] [状态码 4字节] [ ...data... ]
	//[FCS][* CR]
	recv := helper.FromHex(payload[15 : n-4])

	//记录响应的SID
	//t.frame.SID = FromHex(payload[13:15])[0]

	return recv, nil
}


func (t *FinsHostLink) Read(address string, length int) ([]byte, error) {

	//构建读命令
	buf, e := buildReadCommand(address, length)
	if e != nil {
		return nil, e
	}

	//打包命令
	cmd := packAsciiCommand(&t.frame, buf)

	//发送请求
	recv, err := t.request(cmd)
	if err != nil {
		return nil, err
	}

	//[命令码 1 1] [结束码 0 0] , data
	code := helper.ParseUint16(recv[2:])
	if code != 0 {
		return nil, errors.New(fmt.Sprintf("错误码: %d", code))
	}

	return recv[4:], nil
}

func (t *FinsHostLink) Write(address string, values []byte) error {
	//构建写命令
	buf, e := buildWriteCommand(address, values)
	if e != nil {
		return e
	}

	//打包命令
	cmd := packAsciiCommand(&t.frame, buf)

	//发送请求
	recv, err := t.request(cmd)
	if err != nil {
		return err
	}

	//[命令码 1 1] [结束码 0 0]
	code := helper.ParseUint16(recv[2:])
	if code != 0 {
		return errors.New(fmt.Sprintf("错误码: %d", code))
	}

	return nil
}
