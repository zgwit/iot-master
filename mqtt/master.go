package mqtt

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

func subscribeMaster() error {

	//注册应用
	Client.Subscribe("master/register", 0, func(client paho.Client, message paho.Message) {
		var app model.App
		err := json.Unmarshal(message.Payload(), &app)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("app register ", app.Id, " ", app.Name, " ", app.Type, " ", app.Address)
		internal.Applications.Store(app.Id, &app)
	})

	//反注册
	Client.Subscribe("master/unregister", 0, func(client paho.Client, message paho.Message) {
		id := string(message.Payload())
		internal.Applications.Delete(id)
	})

	return nil
}
