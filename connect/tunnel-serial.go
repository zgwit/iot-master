package connect

import (
	"errors"
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/model"
	"time"
)

//TunnelSerial 串口
type TunnelSerial struct {
	tunnelBase
}

func newTunnelSerial(tunnel *model.Tunnel) *TunnelSerial {
	return &TunnelSerial{
		tunnelBase: tunnelBase{tunnel: tunnel},
	}
}

//Open 打开
func (s *TunnelSerial) Open() error {
	if s.running {
		return errors.New("serial is opened")
	}
	s.Emit("open")

	options := serial.OpenOptions{
		PortName:        s.tunnel.Addr,
		BaudRate:        s.tunnel.Serial.BaudRate,
		DataBits:        s.tunnel.Serial.DataBits,
		StopBits:        s.tunnel.Serial.StopBits,
		ParityMode:      serial.ParityMode(s.tunnel.Serial.ParityMode),
		Rs485Enable:     s.tunnel.Serial.RS485,
		MinimumReadSize: 4, //避免只读了一个字节就返回
	}
	port, err := serial.Open(options)
	if err != nil {
		//TODO 串口重试
		return err
	}
	s.running = true
	s.link = port

	//清空重连计数
	s.retry = 0

	go s.receive()

	//上线
	s.tunnel.Last = time.Now()
	_, _ = db.Engine.ID(s.tunnel.Id).Cols("last").Update(s.tunnel)

	return nil
}

//Write 写
func (s *TunnelSerial) Write(data []byte) error {
	if s.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := s.link.Write(data)
	if err != nil {
		s.onClose()
	}
	return err
}

func (s *TunnelSerial) receive() {
	s.running = true
	buf := make([]byte, 1024)
	for {
		n, err := s.link.Read(buf)
		if err != nil {
			s.onClose()
			break
		}
		if n == 0 {
			continue
		}
		//透传转发
		if s.pipe != nil {
			_, err = s.pipe.Write(buf[:n])
			if err != nil {
				s.pipe = nil
			} else {
				continue
			}
		}
		s.Emit("data", buf[:n])
	}
	s.running = false

	retry := &s.tunnel.Retry
	if retry.Enable && (retry.Maximum == 0 || s.retry < retry.Maximum) {
		s.retry++
		time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			_ = s.Open()
		})
	}
}

//Close 关闭
func (s *TunnelSerial) Close() error {
	s.Emit("close")
	if s.link != nil {
		link := s.link
		s.link = nil
		return link.Close()
	}
	s.running = false
	return errors.New("tunnel is closed")
}
