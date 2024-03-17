package device

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/history"
	"github.com/zgwit/iot-master/v4/payload"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"strings"
	"time"
)

func subscribeProperty() {
	mqtt.SubscribeStruct[map[string]any]("device/+/property", func(topic string, values *map[string]any) {
		topics := strings.Split(topic, "/")
		id := topics[1]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			//TODO 自动创建设备？
			return
		}

		//TODO 此处需要判断是 产品 的属性
		//for k, v := range values {
		//	dev.values[k] = v
		//}

		dev.Push(*values)
		dev.Online()

		//写入历史
		err = history.Write(dev.product.Id, dev.id, time.Now().UnixMilli(), *values)
		if err != nil {
			log.Error(err)
		}

	})
}

func mergeProperties(id string, properties []payload.Property) {
	dev, err := Ensure(id)
	if err != nil {
		log.Error(err)
		return
	}

	//合并数据
	for _, p := range properties {
		dev.Values()[p.Name] = p.Value
	}
	dev.Online()
}

func SubscribePropertyStrict() error {
	mqtt.Client.Subscribe("up/property/+/strict", 0, func(client paho.Client, message paho.Message) {
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
