package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
	"github.com/zgwit/iot-master/v3/db"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/vconn"
	"net"
	"net/url"
	"xorm.io/xorm"
)

var Server *mqtt.Server
var Client paho.Client

func Open() error {

	//创建内部Broker
	Server = mqtt.New(nil)

	//TODO 鉴权
	_ = Server.AddHook(new(auth.AllowHook), nil)

	err := mqttCreatePluginListener()
	if err != nil {
		return err
	}

	err = mqttLoadListeners()
	if err != nil {
		return err
	}

	err = Server.Serve()
	if err != nil {
		return err
	}

	err = mqttCreateInternalClient()
	if err != nil {
		return err
	}

	return nil
}

func mqttLoadListeners() error {
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

func mqttCreatePluginListener() error {
	l := listeners.NewTCP("tcp", ":1843", nil)
	err := Server.AddListener(l)
	if err != nil {
		return err
	}

	//unixSock := path.Join(os.TempDir(), "iot-master.sock") //改为临时目录，Windows下兼容性不好，url.Parse错误
	ll := listeners.NewUnixSock("unix", "iot-master.sock")
	return Server.AddListener(ll)
}

func mqttCreateInternalClient() error {
	//client := Server.NewClient(nil, "internal", "internal", true)
	opts := paho.NewClientOptions()
	opts.AddBroker(":1883")
	opts.SetClientID("internal")

	//使用虚拟连接
	opts.SetCustomOpenConnectionFn(func(uri *url.URL, options paho.ClientOptions) (net.Conn, error) {
		c1, c2 := vconn.New()
		//EstablishConnection会读取connect，导致拥堵
		go func() {
			err := Server.EstablishConnection("internal", c1)
			if err != nil {
				log.Error(err)
			}
		}()
		return c2, nil
	})
	// 这里不生效，没搞懂为啥，所以使用SetCustomOpenConnectionFn
	opts.SetDialer(&net.Dialer{
		Resolver: &net.Resolver{Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			c1, c2 := vconn.New()
			_ = Server.EstablishConnection("internal", c1)
			return c2, nil
		}},
	})

	Client = paho.NewClient(opts)
	token := Client.Connect()
	token.Wait()
	err := token.Error()
	if err != nil {
		return err
	}
	//fmt.Println(token.Error())

	return nil
}

func Publish(topic string, payload []byte) error {
	return Server.Publish(topic, payload, false, 0)
}

func PublishJSON(topic string, payload any) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return Server.Publish(topic, bytes, false, 0)
}
