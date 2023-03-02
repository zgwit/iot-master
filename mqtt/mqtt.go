package mqtt

import (
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

func Close() {
	if Server != nil {
		err := Server.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	Client.Disconnect(0)
}

func Open(cfg Options) error {

	var err error
	if cfg.Url != "" {
		//如果指定了外部Broker，则不再开启内置Server
		err = createClient(&cfg)
		if err != nil {
			return err
		}
	} else {
		err = createServer(cfg.Listeners)
		if err != nil {
			return err
		}
		//内部连接
		err = createInternalClient()
		if err != nil {
			return err
		}
	}

	//监听属性
	err = subscribeProperty()
	if err != nil {
		return err
	}

	//监听属性
	err = subscribeMaster()
	if err != nil {
		return err
	}

	return nil
}

func createServer(ls []MqttListener) error {

	//创建内部Broker
	Server = mqtt.New(nil)

	//TODO 鉴权
	_ = Server.AddHook(new(auth.AllowHook), nil)

	err := createListeners(ls)
	if err != nil {
		return err
	}

	err = loadListeners()
	if err != nil {
		return err
	}

	return Server.Serve()
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

func createListeners(ls []MqttListener) error {
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

func createClient(cfg *Options) error {
	//client := Server.NewClient(nil, "internal", "internal", true)
	opts := paho.NewClientOptions()
	opts.AddBroker(cfg.Url)
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)

	Client = paho.NewClient(opts)
	token := Client.Connect()
	token.Wait()
	return token.Error()
}

func createInternalClient() error {
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

func Publish(topic string, payload []byte, retain bool, qos byte) error {
	if Server != nil {
		return Server.Publish(topic, payload, retain, qos)
	}
	//是否需要等待？
	token := Client.Publish(topic, qos, retain, payload)
	_ = token.Wait()
	return token.Error()
}

func PublishJson(topic string, payload any) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return Publish(topic, bytes, false, 0)
}
