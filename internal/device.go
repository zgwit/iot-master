package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
)

var Gateways Map[Device]
var Devices Map[Device]

type Gateway struct {
	Id         string
	Online     bool
	Properties map[string]any
}

type Device struct {
	Id         string
	Properties map[string]any
	//Status map[string]any
}

func subscribeProperty() error {
	MqttClient.Subscribe("up/gateway/+/property", 0, func(client paho.Client, message paho.Message) {
		var prop model.PayloadPropertyUp
		err := json.Unmarshal(message.Payload(), &prop)
		if err != nil {
			return
		}

		//属性值
		if prop.Properties != nil {
			gw := Gateways.Load(prop.Id)
			if gw != nil {
				for _, p := range prop.Properties {
					gw.Properties[p.Name] = p.Value
				}
			}
		}

		//子设备属性
		if prop.Devices != nil {
			for _, d := range prop.Devices {
				dev := Devices.Load(d.Id)
				if dev != nil {
					for _, p := range d.Properties {
						dev.Properties[p.Name] = p.Value
					}
				}
			}
		}

	})

	return nil
}
