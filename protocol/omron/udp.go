package omron

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/protocol/helper"
	"github.com/zgwit/iot-master/service"
)

type UdpFrame struct {
	// 信息控制字段，默认0x80
	ICF byte // 0x80

	// 系统使用的内部信息
	RSV byte // 0x00

	// 网络层信息，默认0x02，如果有八层消息，就设置为0x07
	GCT byte // 0x02

	// PLC的网络号地址，默认0x00
	DNA byte // 0x00

	// PLC的节点地址，这个值在配置了ip地址之后是默认赋值的，默认为Ip地址的最后一位
	DA1 byte // 0x13

	// PLC的单元号地址
	DA2 byte // 0x00

	// 上位机的网络号地址
	SNA byte // 0x00

	// 上位机的节点地址，假如你的电脑的Ip地址为192.168.0.13，那么这个值就是13
	SA1 byte

	// 上位机的单元号地址
	SA2 byte

	// 设备的标识号
	SID byte // 0x00
}

type FinsUdp struct {
	frame UdpFrame
	link  service.Conn
}

func NewFinsUdp(link service.Conn) *FinsUdp {
	return &FinsUdp{
		link: link,
	}
}

func (t *FinsUdp) request(cmd []byte) ([]byte, error) {
	if e := t.link.Write(cmd); e != nil {
		return nil, e
	}

	payload := make([]byte, 1024)
	//t.link.Read(payload)

	//[UDP 10字节]

	//记录响应的SID
	t.frame.SID = payload[9]

	return payload[10:], nil
}

func (t *FinsUdp) Read(address string, length int) ([]byte, error) {

	//构建读命令
	buf, e := buildReadCommand(address, length)
	if e != nil {
		return nil, e
	}

	//打包命令
	cmd := packUDPCommand(&t.frame, buf)

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

func (t *FinsUdp) Write(address string, values []byte) error {
	//构建写命令
	buf, e := buildWriteCommand(address, values)
	if e != nil {
		return e
	}

	//打包命令
	cmd := packUDPCommand(&t.frame, buf)

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
