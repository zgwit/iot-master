package connect

import (
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"sync"
	"time"
)

//NewTunnel 创建通道
func NewTunnel(tunnel *model.Tunnel) (Tunnel, error) {
	var tnl Tunnel
	switch tunnel.Type {
	case "tcp-client":
		tnl = newNetClient(tunnel, "tcp")
		break
	case "tcp-server":
		tnl = newTcpServer(tunnel)
		break
	case "udp-client":
		tnl = newNetClient(tunnel, "udp")
		break
	case "udp-server":
		tnl = newUdpServer(tunnel)
		break
	case "serial":
		tnl = newSerial(tunnel)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", tunnel.Type)
	}

	tnl.On("open", func() {
		_ = database.History.Save(model.TunnelHistory{
			TunnelId: tunnel.Id,
			History:  "open",
			Created:  time.Now(),
		})
	})

	tnl.On("close", func() {
		_ = database.History.Save(model.TunnelHistory{
			TunnelId: tunnel.Id,
			History:  "close",
			Created:  time.Now(),
		})
	})

	tnl.On("link", func(conn Link) {
		_ = database.History.Save(model.LinkHistory{
			LinkId:  conn.ID(),
			History: "online",
			Created: time.Now(),
		})
		conn.Once("close", func() {
			_ = database.History.Save(model.LinkHistory{
				LinkId:  conn.ID(),
				History: "offline",
				Created: time.Now(),
			})
		})
	})

	return tnl, nil
}

var allTunnels sync.Map

//LoadTunnels 加载通道
func LoadTunnels() error {
	var tunnels []*model.Tunnel
	err := database.Master.All(&tunnels)
	if err == storm.ErrNotFound {
		return nil
	}
	for _, t := range tunnels {
		tnl, err := NewTunnel(t)
		if err != nil {
			allTunnels.Store(t.Id, tnl)
		}
	}
	return nil
}
