package device

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/payload"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
	"time"
)

func SubscribeProperty() error {
	mqtt.Client.Subscribe("up/property/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[2]
		id := topics[3]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			//TODO 自动创建设备？
			return
		}

		var values map[string]interface{}
		err = json.Unmarshal(message.Payload(), &values)
		if err != nil {
			return
		}

		//TODO 此处需要判断是 产品 的属性
		//for k, v := range values {
		//	dev.Values[k] = v
		//}
		dev.Push(values)

		dev.Online = true
		dev.Values["$online"] = true
		dev.Values["$update"] = time.Now()
	})

	return nil
}

func mergeProperties(id string, properties []payload.Property) {
	dev, err := Ensure(id)
	if err != nil {
		log.Error(err)
		return
	}

	//合并数据
	for _, p := range properties {
		dev.Values[p.Name] = p.Value
	}
	dev.Values["$online"] = true
	dev.Values["$update"] = time.Now()

}

func SubscribePropertyStrict() error {
	mqtt.Client.Subscribe("up/property/+/+/strict", 0, func(client paho.Client, message paho.Message) {
		var up payload.DevicePropertyUp
		err := json.Unmarshal(message.Payload(), &up)
		if err != nil {
			return
		}

		//属性值
		if up.Id != "" && up.Properties != nil {
			mergeProperties(up.Id, up.Properties)
		}

		//子设备属性
		if up.Devices != nil {
			for _, d := range up.Devices {
				mergeProperties(d.Id, d.Properties)
			}
		}

	})

	return nil
}
