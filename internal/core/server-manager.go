package core

import (
	"fmt"
	"github.com/timshannon/bolthold"
	"github.com/zgwit/iot-master/internal/connect"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/internal/log"
	"github.com/zgwit/iot-master/internal/mqtt"
	"github.com/zgwit/iot-master/link"
	"github.com/zgwit/iot-master/model"
	"sync"
	"time"
)

var allServers sync.Map

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

	svr.On("tunnel", func(tunnel link.Tunnel) {
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
