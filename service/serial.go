package service

import (
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/internal"
	events2 "github.com/zgwit/iot-master/internal/events"
	"time"
)

type Serial struct {
	events2.EventEmitter

	service *internal.Service

	link *SerialLink
}

func newSerial(service *internal.Service) *Serial {
	return &Serial{
		service: service,
	}
}

func (s *Serial) Open() error {
	options := serial.OpenOptions{
		PortName: s.service.Addr,
	}
	if s.service.Serial != nil {
		options.BaudRate = s.service.Serial.BaudRate
		options.DataBits = s.service.Serial.DataBits
		options.StopBits = s.service.Serial.StopBits
		options.ParityMode = serial.ParityMode(s.service.Serial.ParityMode)
	}
	port, err := serial.Open(options)
	if err != nil {
		return err
	}

	s.link = newSerialLink(port)
	go s.link.receive()

	s.Emit("link", s.link)

	//断线后，要重连
	s.link.Once("close", func() {
		retry := s.service.Retry
		if retry == 0 {
			retry = 5 //默认5秒重试
		}
		time.AfterFunc(time.Second*time.Duration(retry), func() {
			_ = s.Open()
		})
	})

	return nil
}

func (s *Serial) Close() error {
	if s.link != nil {
		return s.link.Close()
	}
	return nil //TODO return error
}

func (s *Serial) GetLink(id int) (Link, error) {
	return s.link, nil
}
