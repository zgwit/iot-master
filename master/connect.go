package master

import (
	"fmt"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/protocols"
	"golang.org/x/tools/container/intsets"
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
		has, err := db.Engine.ID(link.Id()).Get(&lnk)
		if err != nil {
			log.Error(err)
			return
		}
		if !has {
			log.Errorf("连接找不到 %d", link.Id())
			return
		}

		//加载协议
		var adapter protocol.Adapter
		adapter, err = protocols.Create(link, tunnel.Protocol.Name, tunnel.Protocol.Options)
		if err != nil {
			log.Error(err)
			//return 无协议，也应该保存起来，只是设备无法正常工作
		}

		allLinks.Store(link.Id(), &Link{Link: lnk, Instance: link, adapter: adapter})

		//第一次连接，初始化默认设备
		if link.First() && tunnel.Devices != nil {
			for _, d := range tunnel.Devices {
				dev :=  model.Device{
					LinkId: link.Id(),
					Station: d.Station,
					ElementId: d.ElementId,
				}
				_, err = db.Engine.InsertOne(&dev)
				if err != nil {
					log.Error(err)
				}
				_, err = LoadDevice(dev.Id)
				if err != nil {
					log.Error(err)
				}
			}
		}

		//找到相关Device，导入Mapper
		var devices []model.Device
		err = db.Engine.Where("link_id", lnk.Id).Find(&devices)
		if err != nil {
			log.Error(err)
			return
		}
		for _, d := range devices {
			dev := GetDevice(d.Id)
			if dev != nil {
				err := dev.Start()
				if err != nil {
					log.Error(err)
					//return
				}
			}
		}

		//连接关闭时，关闭设备
		link.On("close", func() {
			for _, d := range devices {
				dev := GetDevice(d.Id)
				if dev != nil {
					err := dev.Stop()
					if err != nil {
						log.Error(err)
						//return
					}
				}
			}
		})

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
	err := db.Engine.Limit(intsets.MaxInt).Find(&tunnels)
	if err != nil {
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
func LoadTunnel(id int64) error {
	var tunnel model.Tunnel
	has, err := db.Engine.ID(id).Get(&tunnel)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("连接不存在 %d", id)
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
func GetTunnel(id int64) *Tunnel {
	d, ok := allTunnels.Load(id)
	if ok {
		return d.(*Tunnel)
	}
	return nil
}

func RemoveTunnel(id int64) error {
	d, ok := allTunnels.LoadAndDelete(id)
	if ok {
		tnl := d.(*Tunnel)
		return tnl.Instance.Close()
	}
	return nil //error
}

func GetLink(id int64) *Link {
	d, ok := allLinks.Load(id)
	if ok {
		return d.(*Link)
	}
	return nil
}

func RemoveLink(id int64) error {
	d, ok := allLinks.LoadAndDelete(id)
	if ok {
		lnk := d.(*Link)
		return lnk.Instance.Close()
	}
	return nil //error
}