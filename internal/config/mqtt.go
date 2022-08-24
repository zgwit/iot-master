package config

//MQTT 参数
type MQTT struct {
	Addr string
	Sock string
}

var MQTTDefault = MQTT{
	Addr: ":1883",
	Sock: "/iot-master-mqtt.sock",
}
