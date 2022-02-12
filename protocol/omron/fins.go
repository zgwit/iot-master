package omron

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/protocol/helper"
	"github.com/zgwit/iot-master/service"
)

type Fins struct {
	frame UdpFrame
	link  service.Link
}

func NewFins(link service.Link) *Fins {
	return &Fins{
		link: link,
	}
}

func (t *Fins) request(cmd []byte) ([]byte, error) {
	if e := t.link.Write(cmd); e != nil {
		return nil, e
	}

	//接收头16字节：FINS + 长度 + 命令 + 错误码
	buf := make([]byte, 16)
	//_, e := t.link.Read(buf)
	//if e != nil {
	//	return nil, e
	//}

	status := helper.ParseUint32(buf[12:])
	if status != 0 {
		return nil, errors.New(fmt.Sprintf("TCP状态错误: %d", status))
	}

	length := helper.ParseUint32(buf[4:]) - 8

	payload := make([]byte, length)
	//t.link.Read(payload)

	return payload, nil
}

func (t *Fins) Handshake() error {

	// 节点号
	handshake := []byte{0x00, 0x00, 0x00, 0x01}

	cmd := packTCPCommand(0, handshake)

	//发送请求
	buf, e := t.request(cmd)
	if e != nil {
		return e
	}

	//0x00, 0x00, 0x00, 0x01, // 客户端节点号
	//0x00, 0x00, 0x00, 0x01, // PLC端节点号

	//客户端节点
	//t.SA1 = buf[3]
	//服务端节点
	t.frame.DA1 = buf[7]

	return nil
}

func (t *Fins) Read(address string, length int) ([]byte, error) {

	//构建读命令
	buf, e := buildReadCommand(address, length)
	if e != nil {
		return nil, e
	}

	//打包命令
	cmd := packTCPCommand(2, packUDPCommand(&t.frame, buf))

	//发送请求
	recv, err := t.request(cmd)
	if err != nil {
		return nil, err
	}

	//[UDP 10字节] [命令码 1 1] [结束码 0 0] , data

	code := helper.ParseUint16(recv[12:])
	if code != 0 {
		return nil, errors.New(fmt.Sprintf("错误码: %d", code))
	}

	//记录响应的SID
	t.frame.SID = recv[9]

	return recv[14:], nil
}

func (t *Fins) Write(address string, values []byte) error {
	//构建写命令
	buf, e := buildWriteCommand(address, values)
	if e != nil {
		return e
	}

	//打包命令
	cmd := packTCPCommand(2, packUDPCommand(&t.frame, buf))

	//发送请求
	recv, err := t.request(cmd)
	if err != nil {
		return err
	}

	//[UDP 10字节] [命令码 1 1] [结束码 0 0]
	code := helper.ParseUint32(recv[12:])
	if code != 0 {
		return errors.New(fmt.Sprintf("错误码: %d", code))
	}

	//记录响应的SID
	t.frame.SID = recv[9]

	return nil
}
