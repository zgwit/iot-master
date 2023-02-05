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
	if mqttClient != nil {
		mqttClient.Disconnect(0)
	}
	if mqttServer != nil {
		_ = mqttServer.Close()
	}
	//TODO clear gateways devices cache

}
