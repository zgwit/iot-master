package master

import (
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/protocols"
	"github.com/zgwit/storm/v3"
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

func startTunnel(tunnel *model.Tunnel) error {
	tnl, err := connect.NewTunnel(tunnel)
	if err != nil {
		//log.Error(err)
		return err
	}
	allTunnels.Store(tunnel.Id, &Tunnel{
		Tunnel:   *tunnel,
		Instance: tnl,
	})

	tnl.On("link", func(link connect.Link) {
		var lnk model.Link
		err := database.Master.One("Id", link.Id(), &lnk)
		if err != nil && err != storm.ErrNotFound {
			return
		}

		//加载协议
		var adapter protocol.Adapter
		if tunnel.Protocol != nil {
			adapter, err = protocols.Create(link, tunnel.Protocol.Name, tunnel.Protocol.Options)
			if err != nil {
				log.Error(err)
				return
			}
		}

		allLinks.Store(link.Id(), &Link{Link: lnk, Instance: link, adapter: adapter})

		//找到相关Device，导入Mapper
		var devices []model.Device
		err = database.Master.Find("LinkId", link.Id(), &devices)
		if err != nil && err != storm.ErrNotFound {
			return
		}
		for _, d := range devices {
			dev := GetDevice(d.Id)
			if dev != nil {
				err := dev.initMapper()
				if err != nil {
					log.Error(err)
					//return
				}
			}
		}

	})

	err = tnl.Open()
	if err != nil {
		//log.Error(err)
		return err
	}

	return nil
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
	for _, tunnel := range tunnels {
		if tunnel.Disabled {
			continue
		}
		err := startTunnel(tunnel)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

//LoadTunnel 加载通道
func LoadTunnel(id int) error {
	var tunnel model.Tunnel
	err := database.Master.One("Id", id, &tunnel)
	if err != nil {
		return err
	}

	if tunnel.Disabled {
		return nil //TODO error ??
	}
	err = startTunnel(&tunnel)
	if err != nil {
		return err
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
