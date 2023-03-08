package broker

import (
	"fmt"
	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"xorm.io/xorm"
)

var Server *mqtt.Server

func Open(cfg Options) error {

	//创建内部Broker
	Server = mqtt.New(nil)

	//TODO 鉴权
	_ = Server.AddHook(new(auth.AllowHook), nil)

	err := createListeners(cfg.Listeners)
	if err != nil {
		return err
	}

	err = loadListeners()
	if err != nil {
		return err
	}

	return Server.Serve()
}

func Close() {
	if Server != nil {
		err := Server.Close()
		if err != nil {
			log.Error(err)
		}
		Server = nil
	}
}

func loadListeners() error {
	//监听服务
	//加载数据库中 entrypoint
	var entries []model.Server
	err := db.Engine.Find(&entries)
	if err != nil && err != xorm.ErrNotExist {
		return err
	}

	for _, e := range entries {
		id := fmt.Sprintf("tcp-%d", e.Id)
		port := fmt.Sprintf(":%d", e.Port)
		l := listeners.NewTCP(id, port, nil)
		err = Server.AddListener(l)
		if err != nil {
			//return err
			log.Error(err)
		}
	}

	return nil
}

func createListeners(ls []Listener) error {
	for k, l := range ls {
		var err error
		id := fmt.Sprintf("embed-%s-%d", l.Type, k)
		if l.Type == "tcp" {
			err = Server.AddListener(listeners.NewTCP(id, l.Addr, nil))
		} else if l.Type == "unix" {
			err = Server.AddListener(listeners.NewUnixSock(id, l.Addr))
		} else if l.Type == "ws" {
			err = Server.AddListener(listeners.NewWebsocket(id, l.Addr, nil))
		} else {
			return fmt.Errorf("unsupport type %s", l.Type)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
