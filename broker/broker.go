package broker

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	mochi "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/pkg/vconn"
	"net"
	"net/url"
)

var Server *mochi.Server

func Startup() error {
	var brokers []*Broker
	err := db.Engine.Find(&brokers)
	if err != nil {
		return err
	}

	//没有端口
	//if len(brokers) == 0 {
	//	return nil
	//}

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

func CustomOpenConnectionFn(uri *url.URL, options paho.ClientOptions) (net.Conn, error) {
	c1, c2 := vconn.New()
	//EstablishConnection会读取connect，导致拥堵
	go func() {
		err := Server.EstablishConnection("internal", c1)
		if err != nil {
			log.Error(err)
		}
	}()
	return c2, nil
}
