package mqtt

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
	"github.com/zgwit/iot-master/internal/config"
	"github.com/zgwit/iot-master/pkg/vconn"
	"net"
)

var mqttServer *server.Server

func Open(cfg config.MQTT) error {
	mqttServer = server.New()
	mqttServer.Events.OnProcessMessage = func(client events.Client, packet events.Packet) (pkg events.Packet, err error) {

		return
	}

	c := &listeners.Config{
		Auth: new(auth.Allow), //TODO check plugin, mqtt device
	}

	if cfg.Addr != "" {
		l := listeners.NewTCP("tcp", cfg.Addr)
		err := mqttServer.AddListener(l, c)
		if err != nil {
			return err
		}
	}

	if cfg.Sock != "" {
		l := NewUnixSock("unix", cfg.Sock)
		err := mqttServer.AddListener(l, c)
		if err != nil {
			return err
		}
	}

	//TODO websocket

	//物联大师 主连接
	opts := mqtt.NewClientOptions()
	opts.SetDialer(&net.Dialer{
		Resolver: &net.Resolver{Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			c1, c2 := vconn.New()
			_ = mqttServer.EstablishConnection("internal", c1, &auth.Allow{})
			return c2, nil
		}},
	})
	mqtt.NewClient(opts)

	return mqttServer.Serve()
}

func Close() {
	_ = mqttServer.Close()
}

func Publish(topic string, payload []byte) error {
	return mqttServer.Publish(topic, payload, false)
}
