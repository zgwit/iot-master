package master

import (
	"github.com/zgwit/storm/v3"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/protocols"
	"sync"
)

var allTunnels sync.Map
var allLinks sync.Map

type Tunnel struct {
	model.Tunnel
	Instance connect.Tunnel
}

type Link struct {
	model.Link
	Instance connect.Link
	adapter  protocol.Adapter
}

//LoadTunnels 加载通道
func LoadTunnels() error {
	var tunnels []*model.Tunnel
	err := database.Master.All(&tunnels)
	if err == storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}
	for _, d := range tunnels {
		if d.Disabled {
			continue
		}

		tnl, err := connect.NewTunnel(d)
		if err != nil {
			//return err
			//TODO log
			continue
		}
		allTunnels.Store(d.ID, &Tunnel{
			Tunnel:   *d,
			Instance: tnl,
		})

		err = tnl.Open()
		if err != nil {
			//TODO log
		}

		tnl.On("link", func(link connect.Link) {
			var lnk model.Link
			err := database.Master.One("ID", link.ID(), &lnk)
			if err != nil && err != storm.ErrNotFound {
				return
			}

			//加载协议
			var adapter protocol.Adapter
			if lnk.Protocol != nil {
				adapter, err = protocols.Create(link, lnk.Protocol.Name, lnk.Protocol.Options)
				if err != nil {
					//TODO log
					return
				}
			}

			allLinks.Store(link.ID(), &Link{Link: lnk, Instance: link, adapter: adapter})

			//找到相关Device，导入Mapper
			var devices []model.Device
			err = database.Master.Find("LinkID", link.ID(), &devices)
			if err != nil && err != storm.ErrNotFound {
				return
			}
			for _, d := range devices {
				dev := GetDevice(d.ID)
				if dev != nil {
					err := dev.initMapper()
					if err != nil {
						//TODO log
						//return
					}
				}
			}

		})
	}
	return nil
}

//GetTunnel 获取通道
func GetTunnel(id int) *Tunnel {
	d, ok := allTunnels.Load(id)
	if ok {
		return d.(*Tunnel)
	}
	return nil
}

func GetLink(id int) *Link {
	d, ok := allLinks.Load(id)
	if ok {
		return d.(*Link)
	}
	return nil
}
