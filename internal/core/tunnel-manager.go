package core

import (
	"fmt"
	"github.com/timshannon/bolthold"
	"iot-master/internal/connect"
	"iot-master/internal/db"
	"iot-master/internal/log"
	"iot-master/internal/mqtt"
	"iot-master/link"
	"iot-master/model"
	"iot-master/protocols"
	"iot-master/protocols/protocol"
	"sync"
)

var allTunnels sync.Map

type Server struct {
	model.Server
	Instance connect.Server
}

type Tunnel struct {
	model.Tunnel
	Instance link.Tunnel
	protocol protocol.Protocol
}

func bindTunnel(instance link.Tunnel) error {
	tunnel := &Tunnel{
		Tunnel:   *instance.Model(),
		Instance: instance,
		protocol: nil,
	}
	allTunnels.Store(tunnel.Id, tunnel)

	//加载协议
	adapter, err := protocols.Create(instance, tunnel.Protocol.Name, tunnel.Protocol.Options)
	if err != nil {
		return err
	}
	tunnel.protocol = adapter

	//找到相关Device，导入Mapper
	var devices []model.Device
	//err = db.Engine.Where("tunnel_id=?", tunnel.Id).Find(&devices)
	err = db.Store().Find(&devices, bolthold.Where("TunnelId").Eq(tunnel.Id))
	if err != nil {
		return err
	}
	//for _, d := range devices {
	//	dev := GetDevice(d.Id)
	//	if dev != nil {
	//		err := dev.Start()
	//		if err != nil {
	//			log.Error(err)
	//			//return
	//		}
	//	}
	//}

	instance.On("open", func() {
		//TODO 动态加载设备？？？
		_ = mqtt.Publish(fmt.Sprintf("tunnel/%d/open", tunnel.Id), nil)
	})

	instance.On("close", func() {
		_ = mqtt.Publish(fmt.Sprintf("tunnel/%d/close", tunnel.Id), nil)
	})

	instance.On("online", func() {
		_ = mqtt.Publish(fmt.Sprintf("tunnel/%d/online", tunnel.Id), nil)

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
	})

	instance.On("offline", func() {
		_ = mqtt.Publish(fmt.Sprintf("tunnel/%d/offline", tunnel.Id), nil)

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

	return nil
}

func startTunnel(tunnel *model.Tunnel) error {
	tnl, err := connect.NewTunnel(tunnel)
	if err != nil {
		//log.Error(err)
		return err
	}

	err = bindTunnel(tnl)
	if err != nil {
		return err
	}

	return tnl.Open()
}

//LoadTunnels 加载通道
func LoadTunnels() error {
	return db.Store().ForEach(bolthold.Where("ServerId").Eq(0), func(tunnel *model.Tunnel) error {
		if tunnel.Disabled {
			return nil
		}

		go func() {
			err := startTunnel(tunnel)
			if err != nil {
				log.Error(err)
			}
		}()

		return nil
	})
}

//LoadTunnel 加载通道
func LoadTunnel(id uint64) error {
	var tunnel model.Tunnel
	err := db.Store().Get(id, &tunnel)
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

func GetTunnel(id uint64) *Tunnel {
	d, ok := allTunnels.Load(id)
	if ok {
		return d.(*Tunnel)
	}
	return nil
}

func RemoveTunnel(id uint64) error {
	d, ok := allTunnels.LoadAndDelete(id)
	if ok {
		lnk := d.(*Tunnel)
		return lnk.Instance.Close()
	}
	return nil //error
}
