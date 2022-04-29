package connect

import (
	"errors"
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"time"
)

//Serial 串口
type Serial struct {
	events.EventEmitter

	tunnel *model.Tunnel

	link  *SerialLink
	retry int
}

func newSerial(tunnel *model.Tunnel) *Serial {
	return &Serial{
		tunnel: tunnel,
	}
}

//Open 打开
func (s *Serial) Open() error {
	s.Emit("open")

	options := serial.OpenOptions{
		PortName:   s.tunnel.Addr,
		BaudRate:   s.tunnel.Serial.BaudRate,
		DataBits:   s.tunnel.Serial.DataBits,
		StopBits:   s.tunnel.Serial.StopBits,
		ParityMode: serial.ParityMode(s.tunnel.Serial.ParityMode),
	}
	port, err := serial.Open(options)
	if err != nil {
		return err
	}

	//清空重连计数
	s.retry = 0

	s.link = newSerialLink(port)
	go s.link.receive()

	//Store link
	lnk := model.Link{TunnelId: s.tunnel.Id}
	err = database.Master.One("TunnelId", s.tunnel.Id, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Master.Save(&lnk)
	} else if err != nil {
		return err
	} else {
		//上线
	}
	s.link.id = lnk.Id

	s.Emit("link", s.link)

	//断线后，要重连
	s.link.Once("close", func() {
		retry := s.tunnel.Retry
		if retry.Enable && s.retry < retry.Maximum {
			s.retry++
			time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
				_ = s.Open()
			})
		}
	})

	return nil
}

//Close 关闭
func (s *Serial) Close() error {
	s.Emit("close")
	if s.link != nil {
		link := s.link
		s.link = nil
		return link.Close()
	}
	return errors.New("link is closed")
}

func (s *Serial) GetLink(id int) Link {
	return s.link
}
