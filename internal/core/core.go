package core

func Open() error {

	err := openMqttServer()
	if err != nil {
		return err
	}

	err = subscribeProperty()
	if err != nil {
		return err
	}

	err = subscribeService()
	if err != nil {
		return err
	}

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
