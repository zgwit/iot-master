package service

import (
	"fmt"
	"github.com/zgwit/iot-master/model"
)

type Service interface {
	Open() error
	Close() error
	GetLink(id int)(Link, error)
}

func NewService(service *model.Service) (Service, error)  {
	var svc Service
	switch service.Type {
	case "tcp-client":
		svc = NewTcpClient(service)
		break
	case "tcp-server":
		svc = NewTcpServer(service)
		break
	case "udp-client":
		svc = NewUdpClient(service)
		break
	case "udp-server":
		svc = NewUdpServer(service)
		break
	case "serial":
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", service.Type)
	}
	return svc, nil
}