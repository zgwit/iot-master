package internal

import (
	"context"
	"encoding/json"
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
	"github.com/zgwit/iot-master/v3/internal/db"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/vconn"
	"net"
	"xorm.io/xorm"
)

var mqttServer *mqtt.Server
var mqttClient paho.Client

func openMqttServer() error {

	//创建内部Broker
	mqttServer = mqtt.New(nil)

	//TODO 鉴权
	_ = mqttServer.AddHook(new(auth.AllowHook), nil)

	err := mqttLoadListeners()
	if err != nil {
		return err
	}

	err = mqttCreatePluginListener()
	if err != nil {
		return err
	}

	err = mqttServer.Serve()
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
		err = mqttServer.AddListener(l)
		if err != nil {
			//return err
			log.Error(err)
		}
	}

	return nil
}

func mqttCreatePluginListener() error {
	l := listeners.NewUnixSock("plugin", "iot-master.sock")
	return mqttServer.AddListener(l)
}

func mqttCreateInternalClient() error {
	//client := mqttServer.NewClient(nil, "internal", "internal", true)
	opts := paho.NewClientOptions()
	opts.AddBroker(":1883")
	opts.SetClientID("internal")
	//TODO 这里不生效，为啥
	opts.SetDialer(&net.Dialer{
		Resolver: &net.Resolver{Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			c1, c2 := vconn.New()
			_ = mqttServer.EstablishConnection("internal", c1)
			return c2, nil
		}},
	})

	mqttClient = paho.NewClient(opts)
	token := mqttClient.Connect()
	token.Wait()
	err := token.Error()
	if err != nil {
		return err
	}
	//fmt.Println(token.Error())

	//订阅消息
	//subscribeTopics(MQTT)
	return nil
}

func Publish(topic string, payload []byte) error {
	return mqttServer.Publish(topic, payload, false, 0)
}

func PublishEx(topic string, payload any) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return mqttServer.Publish(topic, bytes, false, 0)
}
