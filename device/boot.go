package device

func Open() error {

	mqttEvent()

	mqttOnline()

	mqttProperty()

	return nil
}

func Close() {
	//TODO 取消MQTT订阅
}
