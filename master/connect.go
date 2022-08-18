package master

import (
	"fmt"
	"github.com/timshannon/bolthold"
	"iot-master/conn"
	"iot-master/internal/connect"
	"iot-master/internal/db"
	"iot-master/internal/log"
	"iot-master/internal/mqtt"
	"iot-master/model"
	"iot-master/protocols"
	"iot-master/protocols/protocol"
	"sync"
	"time"
)

var allServers sync.Map
var allTunnels sync.Map

type Server struct {
	model.Server
	Instance connect.Server
}

type Tunnel struct {
	model.Tunnel
	Instance conn.Tunnel
	protocol protocol.Protocol
}

func bindTunnel(instance conn.Tunnel) error {
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

func startServer(server *model.Server) error {
	svr, err := connect.NewServer(server)
	if err != nil {
		//log.Error(err)
		return err
	}
	allServers.Store(server.Id, &Server{
		Server:   *server,
		Instance: svr,
	})

	svr.On("open", func() {
		//TODO 加载设备？？？
		_ = mqtt.Publish(fmt.Sprintf("server/%d/open", server.Id), nil)
	})

	svr.On("close", func() {
		_ = mqtt.Publish(fmt.Sprintf("server/%d/close", server.Id), nil)
	})

	svr.On("tunnel", func(tunnel conn.Tunnel) {
		//第一次连接，初始化默认设备
		if tunnel.First() && server.Devices != nil {
			for _, d := range server.Devices {
				dev := model.Device{
					TunnelId:  tunnel.Model().Id,
					Station:   d.Station,
					ProductId: d.ProductId,
					Created:   time.Now(),
				}
				err := db.Store().Insert(bolthold.NextSequence(), &dev)
				if err != nil {
					log.Error(err)
				}
				_, err = LoadDevice(dev.Id)
				if err != nil {
					log.Error(err)
				}
			}
		}

		err := bindTunnel(tunnel)
		if err != nil {
			log.Error(err)
			//return 无协议，也应该保存起来，只是设备无法正常工作
		}
	})

	err = svr.Open()
	if err != nil {
		//log.Error(err)
		return err
	}

	return nil
}

//LoadServers 加载通道
func LoadServers() error {
	return db.Store().ForEach(nil, func(server *model.Server) error {
		if server.Disabled {
			return nil
		}
		go func() {
			err := startServer(server)
			if err != nil {
				log.Error(err)
			}
		}()
		return nil
	})
}

//LoadServer 加载通道
func LoadServer(id uint64) error {
	var server model.Server
	err := db.Store().Get(id, &server)
	if err != nil {
		return err
	}
	if server.Disabled {
		return nil //TODO error ??
	}
	err = startServer(&server)
	if err != nil {
		return err
	}
	return nil
}

//GetServer 获取通道
func GetServer(id uint64) *Server {
	d, ok := allServers.Load(id)
	if ok {
		return d.(*Server)
	}
	return nil
}

func RemoveServer(id uint64) error {
	d, ok := allServers.LoadAndDelete(id)
	if ok {
		tnl := d.(*Server)
		return tnl.Instance.Close()
	}
	return nil //error
}
