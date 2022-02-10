package service

import (
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/common"
	"github.com/zgwit/iot-master/model"
)

type Serial struct {
	common.EventEmitter

	service *model.Service

	link *SerialLink
}

func newSerial(service *model.Service) *Serial {
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

	//TODO 断线后，要重连

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
