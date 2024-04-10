package device

import "github.com/zgwit/iot-master/v4/boot"

func init() {
	boot.Register("device", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"database", "mqtt"},
	})
}

func Startup() error {

	mqttEvent()

	mqttOnline()

	mqttProperty()

	return nil
}

func Shutdown() error {
	//TODO 取消MQTT订阅

	return nil
}
