package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
)

var Applications Map[model.App]

func subscribeMaster() error {

	//注册应用
	MqttClient.Subscribe("master/register", 0, func(client paho.Client, message paho.Message) {
		var svc model.App
		err := json.Unmarshal(message.Payload(), &svc)
		if err != nil {
			return
		}
		Applications.Store(svc.Id, &svc)
	})

	//反注册
	MqttClient.Subscribe("master/unregister", 0, func(client paho.Client, message paho.Message) {
		id := string(message.Payload())
		Applications.Delete(id)
	})

	return nil
}
