package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/mqtt"
)

func subscribeProperty() error {
	mqtt.Client.Subscribe("up/property/+/+", 0, func(client paho.Client, message paho.Message) {
		var payload model.UpPropertyPayload
		err := json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			return
		}

		//属性值
		if payload.Id != "" && payload.Properties != nil {
			dev := Devices.Load(payload.Id)
			if dev == nil {
				dev = NewDevice(payload.Id)
				Devices.Store(payload.Id, dev)
			}
			//合并数据
			for _, p := range payload.Properties {
				dev.Properties[p.Name] = p.Value
			}
		}

		//子设备属性
		if payload.Devices != nil {
			for _, d := range payload.Devices {
				dev := Subsets.Load(d.Id)
				if dev == nil {
					dev = NewSubset(payload.Id)
					Subsets.Store(payload.Id, dev)
				}
				//合并数据
				for _, p := range d.Properties {
					dev.Properties[p.Name] = p.Value
				}
			}
		}

	})

	return nil
}
