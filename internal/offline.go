package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func subscribeOffline() error {
	mqtt.Client.Subscribe("offline/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[2]
		id := topics[3]

		dev, err := GetDevice(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = false
		dev.Values["$online"] = false
	})

	return nil
}
