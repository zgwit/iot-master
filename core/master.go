package core

func SubscribeMaster() error {
	//注册应用
	//mqtt.Subscribe[types.App]("master/register", func(topic string, a *types.App) {
	//	log.Info("a register ", a.id, " ", a.name, " ", a.Type, " ", a.Address)
	//	plugin.Applications.Store(a.id, a)
	//})

	//反注册
	//mqtt.Subscribe[any]("master/unregister", func(topic string, payload *any) {
	//	id := string(payload)
	//	app.Applications.Delete(id)
	//})

	return nil
}
