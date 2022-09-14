package broker

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
	"github.com/zgwit/iot-master/pkg/vconn"
	"net"
)

var mqttServer *server.Server
var MQTT mqtt.Client

func Open(cfg Options) error {

	internal := cfg.Url != "" && cfg.Url != "internal"

	//创建内部Broker
	if internal {
		mqttServer = server.New()
		c := &listeners.Config{
			Auth: new(auth.Allow), //TODO check plugin, mqtt device
		}

		//TODO websocket
		l := listeners.NewTCP("tcp", cfg.Url)
		err := mqttServer.AddListener(l, c)
		if err != nil {
			return err
		}

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
