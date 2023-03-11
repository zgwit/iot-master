package internal

func Open() error {

	err := LoadProducts()
	if err != nil {
		return err
	}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	err = subscribeMaster()
	if err != nil {
		return err
	}

	err = subscribeProperty()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	//TODO clear gateways devices data

}
