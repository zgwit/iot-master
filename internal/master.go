package internal

import (
	app2 "github.com/zgwit/iot-master/v4/app"
	"github.com/zgwit/iot-master/v4/model"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
)

func SubscribeMaster() error {
	//注册应用
	mqtt.Subscribe[model.App]("master/register", func(topic string, app *model.App) {
		log.Info("app register ", app.Id, " ", app.Name, " ", app.Type, " ", app.Address)
		app2.Applications.Store(app.Id, app)
	})

	//反注册
	//mqtt.Subscribe[any]("master/unregister", func(topic string, payload *any) {
	//	id := string(payload)
	//	app2.Applications.Delete(id)
	//})

	return nil
}
