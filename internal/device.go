package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/mqtt"
	"time"
)

var Devices Map[Device]

type Device struct {
	Id         string
	Last       time.Time
	Properties map[string]any
}

func subscribeProperty() error {
	mqtt.Client.Subscribe("up/property/+/+", 0, func(client paho.Client, message paho.Message) {
		var prop model.UpPropertyPayload
		err := json.Unmarshal(message.Payload(), &prop)
		if err != nil {
			return
		}

		//属性值
		if prop.Properties != nil {
			dev := Devices.Load(prop.Id)
			if dev != nil {
				for _, p := range prop.Properties {
					dev.Properties[p.Name] = p.Value
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
