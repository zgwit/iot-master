package service

import (
	"fmt"
	"github.com/zgwit/iot-master/internal"
	events2 "github.com/zgwit/iot-master/internal/events"
)

type Service interface {
	events2.EventEmitterInterface

	Open() error
	Close() error
	GetLink(id int)(Link, error)
}

func NewService(service *internal.Service) (Service, error)  {
	var svc Service
	switch service.Type {
	case "tcp-client":
		svc = newNetClient(service, "tcp")
		break
	case "tcp-server":
		svc = newTcpServer(service)
		break
	case "udp-client":
		svc = newNetClient(service, "udp")
		break
	case "udp-server":
		svc = NewUdpServer(service)
		break
	case "serial":
		svc = newSerial(service)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", service.Type)
	}
	return svc, nil
}