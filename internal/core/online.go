package core

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func subscribeOnline() error {
	mqtt.Client.Subscribe("online/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[1]
		id := topics[2]

		dev, err := GetDevice(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = true
		dev.Values["$online"] = true
	})

	mqtt.Client.Subscribe("offline/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[1]
		id := topics[2]

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
