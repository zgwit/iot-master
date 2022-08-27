package config

import (
	"os"
	"path"
)

// MQTT 参数
type MQTT struct {
	Addr string
	Sock string
}

var MQTTDefault = MQTT{
	Addr: ":1883",
	Sock: path.Join(os.TempDir(), "iot-master-mqtt.sock"),
}
