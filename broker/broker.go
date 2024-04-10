package broker

import (
	"fmt"
	mochi "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
)

var Server *mochi.Server

func Startup() error {
	var brokers []*Broker
	err := db.Engine.Find(&brokers)
	if err != nil {
		return err
	}

	//没有端口
	if len(brokers) == 0 {
		return nil
	}

	Server = mochi.New(nil)
	//Server.NewClient()

	for _, e := range brokers {
		l := listeners.NewTCP(listeners.Config{
			Type:    "tcp",
			ID:      fmt.Sprintf("tcp-%s", e.Id),
			Address: fmt.Sprintf(":%d", e.Port),
		})
		err = Server.AddListener(l)
		if err != nil {
			//return err
			log.Error(err)
		}
	}

	//TODO 鉴权
	_ = Server.AddHook(new(Auth), nil)

	//监听默认端口
	//port := config.GetInt(MODULE, "port")
	//addr := ":" + strconv.Itoa(port)
	//err := Server.AddListener(listeners.NewTCP(listeners.Config{
	//	Type:    "tcp",
	//	ID:      "embed",
	//	Address: addr,
	//}))
	//if err != nil {
	//	return err
	//}

	//监听UnixSocket，Win10以下版本有问题
	//if options.Unix {
	//	err = Server.AddListener(listeners.NewUnixSock("embed-unix", options.Addr))
	//	if err != nil {
	//		log.Error(err)
	//		//return err
	//	}
	//}

	return Server.Serve()
}

func Shutdown() (err error) {
	if Server != nil {
		err = Server.Close()
		Server = nil
	}
	return
}

func loadListeners() error {
	//监听服务
	//加载数据库中 entrypoint
	var entries []*Broker
	err := db.Engine.Find(&entries)
	if err != nil {
		return err
	}

	for _, e := range entries {
		l := listeners.NewTCP(listeners.Config{
			Type:    "tcp",
			ID:      fmt.Sprintf("tcp-%s", e.Id),
			Address: fmt.Sprintf(":%d", e.Port),
		})
		err = Server.AddListener(l)
		if err != nil {
			//return err
			log.Error(err)
		}
	}

	return nil
}
