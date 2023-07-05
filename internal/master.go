package internal

import (
	app2 "github.com/zgwit/iot-master/v3/app"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

func SubscribeMaster() error {
	//注册应用
	mqtt.SubscribeStruct[model.App]("master/register", func(topic string, app *model.App) {
		log.Info("app register ", app.Id, " ", app.Name, " ", app.Type, " ", app.Address)
		app2.Applications.Store(app.Id, app)
	})

	//反注册
	mqtt.Subscribe("master/unregister", func(topic string, payload []byte) {
		id := string(payload)
		app2.Applications.Delete(id)
	})

	return nil
}
