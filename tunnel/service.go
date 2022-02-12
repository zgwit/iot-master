package tunnel

import (
	"fmt"
	"github.com/zgwit/iot-master/events"
)

type Service interface {
	events.EventInterface

	Open() error
	Close() error
	GetLink(id int)(Conn, error)
}

func NewService(tunnel *Tunnel) (Service, error)  {
	var svc Service
	switch tunnel.Type {
	case "tcp-client":
		svc = newNetClient(tunnel, "tcp")
		break
	case "tcp-server":
		svc = newTcpServer(tunnel)
		break
	case "udp-client":
		svc = newNetClient(tunnel, "udp")
		break
	case "udp-server":
		svc = NewUdpServer(tunnel)
		break
	case "serial":
		svc = newSerial(tunnel)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", tunnel.Type)
	}
	return svc, nil
}