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

var allServers sync.Map
var allTunnels sync.Map

type Server struct {
	model.Server
	Instance connect.Server
}

type Tunnel struct {
	model.Tunnel
	Instance connect.Tunnel
	adapter  protocol.Adapter
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

	svr.On("tunnel", func(tunnel connect.Tunnel) {
		mod := tunnel.Model()

		//加载协议
		var adapter protocol.Adapter
		adapter, err = protocols.Create(tunnel, mod.Protocol.Name, mod.Protocol.Options)
		if err != nil {
			log.Error(err)
			//return 无协议，也应该保存起来，只是设备无法正常工作
		}

		allTunnels.Store(mod, &Tunnel{Tunnel: *mod, Instance: tunnel, adapter: adapter})

		//第一次连接，初始化默认设备
		if tunnel.First() && server.Devices != nil {
			for _, d := range server.Devices {
				dev := model.Device{
					TunnelId:  mod.Id,
					Station:   d.Station,
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
		err = db.Engine.Where("tunnel_id", mod.Id).Find(&devices)
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
		tunnel.On("close", func() {
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

	err = svr.Open()
	if err != nil {
		//log.Error(err)
		return err
	}

	return nil
}

//LoadServers 加载通道
func LoadServers() error {
	var servers []*model.Server
	err := db.Engine.Limit(intsets.MaxInt).Find(&servers)
	if err != nil {
		return err
	}
	for _, server := range servers {
		if server.Disabled {
			continue
		}

		server := server //避免for闭包问题
		go func() {
			err := startServer(server)
			if err != nil {
				log.Error(err)
			}
		}()
	}
	return nil
}

//LoadServer 加载通道
func LoadServer(id int64) error {
	var tunnel model.Server
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
	err = startServer(&tunnel)
	if err != nil {
		return err
	}
	return nil
}

//GetServer 获取通道
func GetServer(id int64) *Server {
	d, ok := allServers.Load(id)
	if ok {
		return d.(*Server)
	}
	return nil
}

func RemoveServer(id int64) error {
	d, ok := allServers.LoadAndDelete(id)
	if ok {
		tnl := d.(*Server)
		return tnl.Instance.Close()
	}
	return nil //error
}

func startTunnel(tunnel *model.Tunnel) error {
	tnl, err := connect.NewTunnel(tunnel)
	if err != nil {
		//log.Error(err)
		return err
	}

	instance := &Tunnel{
		Tunnel:   *tunnel,
		Instance: tnl,
	}
	allTunnels.Store(tunnel.Id, instance)

	//加载协议
	var adapter protocol.Adapter
	adapter, err = protocols.Create(tnl, tunnel.Protocol.Name, tunnel.Protocol.Options)
	if err != nil {
		log.Error(err)
		//return 无协议，也应该保存起来，只是设备无法正常工作
	}

	instance.adapter = adapter

	//找到相关Device，导入Mapper
	var devices []model.Device
	err = db.Engine.Where("tunnel_id", tunnel.Id).Find(&devices)
	if err != nil {
		return err
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
	tnl.On("close", func() {
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

	err = tnl.Open()
	if err != nil {
		return err
	}
	return nil
}

//LoadTunnels 加载通道
func LoadTunnels() error {
	var tunnels []*model.Tunnel
	err := db.Engine.Limit(intsets.MaxInt).Where("server_id IS NULL").Find(&tunnels)
	if err != nil {
		return err
	}
	for _, tunnel := range tunnels {
		if tunnel.Disabled {
			continue
		}

		tunnel := tunnel //避免range闭包问题
		go func() {
			err := startTunnel(tunnel)
			if err != nil {
				log.Error(err)
			}
		}()
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
		lnk := d.(*Tunnel)
		return lnk.Instance.Close()
	}
	return nil //error
}
