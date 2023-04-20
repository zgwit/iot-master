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

	err := createEmbedListener(cfg)
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
	var entries []model.Broker
	err := db.Engine.Find(&entries)
	if err != nil && err != xorm.ErrNotExist {
		return err
	}

	for _, e := range entries {
		id := fmt.Sprintf("tcp-%s", e.Id)
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

func createEmbedListener(opts Options) (err error) {
	id := fmt.Sprintf("embed-%s", opts.Type)
	if opts.Type == "tcp" {
		err = Server.AddListener(listeners.NewTCP(id, opts.Addr, nil))
	} else if opts.Type == "unix" {
		err = Server.AddListener(listeners.NewUnixSock(id, opts.Addr))
	} else if opts.Type == "websocket" {
		err = Server.AddListener(listeners.NewWebsocket(id, opts.Addr, nil))
	} else {
		err = fmt.Errorf("unsupport type %s", opts.Type)
	}
	return
}
