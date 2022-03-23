package connect

import (
	"errors"
	"github.com/asdine/storm/v3"
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"time"
)

//Serial 串口
type Serial struct {
	events.EventEmitter

	tunnel *model.Tunnel

	link *SerialLink
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
		PortName: s.tunnel.Addr,
	}
	if s.tunnel.Serial != nil {
		options.BaudRate = s.tunnel.Serial.BaudRate
		options.DataBits = s.tunnel.Serial.DataBits
		options.StopBits = s.tunnel.Serial.StopBits
		options.ParityMode = serial.ParityMode(s.tunnel.Serial.ParityMode)
	}
	port, err := serial.Open(options)
	if err != nil {
		return err
	}

	s.link = newSerialLink(port)
	go s.link.receive()
	
	//Store link
	lnk := model.Link{
		TunnelID: s.tunnel.ID,
		Protocol: s.tunnel.Protocol,
		Created:  time.Now(),
	}
	err = database.Master.One("TunnelID", s.tunnel.ID, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Master.Save(&lnk)
	} else if err != nil {
		return err
	} else {
		//上线
	}
	s.link.id = lnk.ID
	
	s.Emit("link", s.link)

	//断线后，要重连
	s.link.Once("close", func() {
		retry := s.tunnel.Retry
		if retry == 0 {
			retry = 5 //默认5秒重试
		}
		time.AfterFunc(time.Second*time.Duration(retry), func() {
			_ = s.Open()
		})
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
