package core

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
)

var Services Map[Service]

type Service struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Type    string `json:"type"`
	Address string `json:"address"`
}

func subscribeService() error {
	mqttClient.Subscribe("$service/register", 0, func(client paho.Client, message paho.Message) {
		var svc Service
		err := json.Unmarshal(message.Payload(), &svc)
		if err != nil {
			return
		}
		Services.Store(svc.Name, &svc)
	})
	return nil
}
