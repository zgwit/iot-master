package connect

import (
	"fmt"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
)

//Server 通道
type Server interface {
	events.EventInterface

	Open() error
	Close() error
	GetTunnel(id int64) Tunnel
	Running() bool
}

//NewServer 创建通道
func NewServer(server *model.Server) (Server, error) {
	var tnl Server
	switch server.Type {
	case "tcp":
		tnl = newServerTCP(server)
		break
	case "udp":
		tnl = newServerUDP(server)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", server.Type)
	}

	tnl.On("open", func() {
		_, _ = db.Engine.InsertOne(model.Event{
			Target:   "server",
			TargetId: server.Id,
			Event:    "打开",
		})
	})

	tnl.On("close", func() {
		_, _ = db.Engine.InsertOne(model.Event{
			Target:   "server",
			TargetId: server.Id,
			Event:    "关闭",
		})
	})

	return tnl, nil
}
