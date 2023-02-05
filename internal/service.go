package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
)

var Applications Map[model.App]

func subscribeService() error {

	//注册应用
	mqttClient.Subscribe("master/register", 0, func(client paho.Client, message paho.Message) {
		var svc model.App
		err := json.Unmarshal(message.Payload(), &svc)
		if err != nil {
			return
		}
		Applications.Store(svc.Name, &svc)
	})

	return nil
}
