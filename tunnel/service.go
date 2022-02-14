package tunnel

import (
	"fmt"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"time"
)

type Service interface {
	events.EventInterface

	Open() error
	Close() error
	GetLink(id int) (Conn, error)
}

func NewService(tunnel *Tunnel) (Service, error) {
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

	svc.On("open", func() {
		_ = database.TunnelHistory.Save(TunnelHistory{
			ServiceId: tunnel.Id,
			History:   "open",
			Created:   time.Now(),
		})
	})

	svc.On("close", func() {
		_ = database.TunnelHistory.Save(TunnelHistory{
			ServiceId: tunnel.Id,
			History:   "close",
			Created:   time.Now(),
		})
	})

	svc.On("link", func(conn Conn) {
		_ = database.LinkHistory.Save(LinkHistory{
			LinkId:  conn.ID(),
			History: "online",
			Created: time.Now(),
		})
		conn.Once("close", func() {
			_ = database.LinkHistory.Save(LinkHistory{
				LinkId:  conn.ID(),
				History: "offline",
				Created: time.Now(),
			})
		})
	})

	return svc, nil
}
