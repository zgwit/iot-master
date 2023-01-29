package core

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
	"github.com/zgwit/iot-master/v3/internal/db"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/vconn"
	"net"
	"xorm.io/xorm"
)

var mqttServer *server.Server
var MQTT mqtt.Client

func Open() error {

	//创建内部Broker
	mqttServer = server.New()

	mqttServer.Events.OnConnect = func(client events.Client, packet events.Packet) {
		//自动创建网关？ 处理连接状态？
	}

	mqttServer.Events.OnDisconnect = func(client events.Client, err error) {
		//网关离线 client.ID
	}

	err := loadListeners()
	if err != nil {
		return err
	}

	//TODO websocket

	err = mqttServer.Serve()
	if err != nil {
		return err
	}

	//物联大师 主连接
	opts := mqtt.NewClientOptions()
	opts.AddBroker(":1883")
	opts.SetClientID("iot-master")
	//TODO 这里不生效，为啥
	opts.SetDialer(&net.Dialer{
		Resolver: &net.Resolver{Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			c1, c2 := vconn.New()
			_ = mqttServer.EstablishConnection("internal", c1, &auth.Allow{})
			return c2, nil
		}},
	})

	MQTT = mqtt.NewClient(opts)
	token := MQTT.Connect()
	token.Wait()
	fmt.Println(token.Error())

	//订阅消息
	subscribeTopics(MQTT)

	return nil
}

func loadListeners() error {
	//监听服务
	c := &listeners.Config{
		Auth: new(auth.Allow), //TODO check plugin, mqtt device
	}

	var entries []model.Entrypoint
	err := db.Engine.Find(&entries)
	if err != nil && err != xorm.ErrNotExist {
		return err
	}

	for _, e := range entries {
		//TODO 加载数据库中 entrypoint
		l := listeners.NewTCP("tcp", fmt.Sprintf(":%d", e.Port))
		err = mqttServer.AddListener(l, c)
		if err != nil {
			//return err
			log.Error(err)
		}
	}

	return nil
}

func Close() {
	if mqttServer != nil {
		_ = mqttServer.Close()
	}
	MQTT.Disconnect(0)
}

func Publish(topic string, payload []byte) error {
	//TODO 兼容struct类型

	if mqttServer != nil {
		return mqttServer.Publish(topic, payload, false)
	}

	MQTT.Publish(topic, 0, false, payload)
	return nil
}
