package mqtt

import (
	"github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
)

var mqttServer *server.Server

func Open() error {
	mqttServer = server.New()

	c := &listeners.Config{
		Auth: new(auth.Allow), //TODO check plugin, mqtt device
	}

	l := listeners.NewTCP("tcp", ":1883")
	err := mqttServer.AddListener(l, c)
	if err != nil {
		return err
	}

	//TODO websocket

	return mqttServer.Serve()
}

func Publish(topic string, payload []byte) error {
	return mqttServer.Publish(topic, payload, false)
}
