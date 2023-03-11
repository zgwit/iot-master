package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

func mergeProperties(id string, properties []model.PayloadValue) {
	dev := Devices.Load(id)
	if dev == nil {
		//加载设备
		err := LoadDeviceById(id)
		if err != nil {
			log.Error(err)
			return
		}

		dev = Devices.Load(id)
	}
	//合并数据
	for _, p := range properties {
		dev.Values[p.Name] = p.Value
	}
}

func subscribeProperty() error {
	mqtt.Client.Subscribe("up/property/+/+", 0, func(client paho.Client, message paho.Message) {
		var payload model.PayloadPropertyUp
		err := json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			return
		}

		//属性值
		if payload.Id != "" && payload.Properties != nil {
			mergeProperties(payload.Id, payload.Properties)
		}

		//子设备属性
		if payload.Devices != nil {
			for _, d := range payload.Devices {
				mergeProperties(d.Id, d.Properties)
			}
		}

	})

	return nil
}
