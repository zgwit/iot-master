package internal

import (
	"github.com/zgwit/iot-master/v4/app"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/types"
)

func SubscribeMaster() error {
	//注册应用
	mqtt.Subscribe[types.App]("master/register", func(topic string, a *types.App) {
		log.Info("a register ", a.Id, " ", a.Name, " ", a.Type, " ", a.Address)
		app.Applications.Store(a.Id, a)
	})

	//反注册
	//mqtt.Subscribe[any]("master/unregister", func(topic string, payload *any) {
	//	id := string(payload)
	//	app.Applications.Delete(id)
	//})

	return nil
}
