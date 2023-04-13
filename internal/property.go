package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
	"time"
)

func subscribeProperty() error {
	mqtt.Client.Subscribe("up/property/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[2]
		id := topics[3]

		dev, err := GetDevice(id)
		if err != nil {
			log.Error(err)
			//TODO 自动创建设备？
			return
		}

		var payload map[string]interface{}
		err = json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			return
		}

		//TODO 此处需要判断是 产品 的属性
		for k, v := range payload {
			dev.Values[k] = v
		}
		dev.Online = true
		dev.Values["$online"] = true
		dev.Values["$update"] = model.Time(time.Now())
	})

	return nil
}

func mergeProperties(id string, properties []model.PayloadValue) {
	dev, err := GetDevice(id)
	if err != nil {
		log.Error(err)
		return
	}

	//合并数据
	for _, p := range properties {
		dev.Values[p.Name] = p.Value
	}
	dev.Values["$online"] = true
	dev.Values["$update"] = model.Time(time.Now())

}

func subscribePropertyStrict() error {
	mqtt.Client.Subscribe("up/property/+/+/strict", 0, func(client paho.Client, message paho.Message) {
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
