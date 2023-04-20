package core

func Open() error {

	err := LoadProducts()
	if err != nil {
		return err
	}

	//err = LoadDevices()
	//if err != nil {
	//	return err
	//}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	err = subscribeMaster()
	if err != nil {
		return err
	}

	err = SubscribeEvent()
	if err != nil {
		return err
	}

	err = subscribeProperty()
	if err != nil {
		return err
	}

	err = subscribePropertyStrict()
	if err != nil {
		return err
	}

	err = subscribeOnline()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	//TODO clear gateways devices data

}
