package core

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
	"github.com/zgwit/iot-master/v2/internal/broker"
	"github.com/zgwit/iot-master/v2/pkg/vconn"
	"net"
)

var mqttServer *server.Server
var MQTT mqtt.Client

func Open(cfg broker.Options) error {

	internal := cfg.Url == "" || cfg.Url == "internal"

	//创建内部Broker
	if internal {
		mqttServer = server.New()
		mqttServer.Events.OnConnect = func(client events.Client, packet events.Packet) {
			//自动创建网关？ 处理连接状态？
		}
		mqttServer.Events.OnDisconnect = func(client events.Client, err error) {
			//网关离线 client.ID
		}

		//监听服务
		c := &listeners.Config{
			Auth: new(auth.Allow), //TODO check plugin, mqtt device
		}

		l := listeners.NewTCP("tcp", ":1883")
		err := mqttServer.AddListener(l, c)
		if err != nil {
			return err
		}

		//TODO websocket

		err = mqttServer.Serve()
		if err != nil {
			return err
		}
	}

	//物联大师 主连接
	opts := mqtt.NewClientOptions()
	if internal {
		opts.SetDialer(&net.Dialer{
			Resolver: &net.Resolver{Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				c1, c2 := vconn.New()
				_ = mqttServer.EstablishConnection("internal", c1, &auth.Allow{})
				return c2, nil
			}},
		})
	} else {
		opts.AddBroker(cfg.Url)
		opts.SetClientID(cfg.ClientId)
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
	}
	MQTT = mqtt.NewClient(opts)

	//订阅消息
	subscribeTopics(MQTT)

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
