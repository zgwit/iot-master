package internal

func Open() error {
	err := openMqttServer()
	if err != nil {
		return err
	}

	err = subscribeProperty()
	if err != nil {
		return err
	}

	err = subscribeMaster()
	if err != nil {
		return err
	}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	return nil
}

func Close() {
	if MqttClient != nil {
		MqttClient.Disconnect(0)
	}
	if MqttServer != nil {
		_ = MqttServer.Close()
	}
	//TODO clear gateways devices cache

}
