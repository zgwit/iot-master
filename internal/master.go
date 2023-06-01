package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	app2 "github.com/zgwit/iot-master/v3/app"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

func SubscribeMaster() error {

	//注册应用
	mqtt.Client.Subscribe("master/register", 0, func(client paho.Client, message paho.Message) {
		var app model.App
		err := json.Unmarshal(message.Payload(), &app)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("app register ", app.Id, " ", app.Name, " ", app.Type, " ", app.Address)
		app2.Applications.Store(app.Id, &app)
	})

	//反注册
	mqtt.Client.Subscribe("master/unregister", 0, func(client paho.Client, message paho.Message) {
		id := string(message.Payload())
		app2.Applications.Delete(id)
	})

	return nil
}
