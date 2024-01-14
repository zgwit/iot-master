package device

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/payload"
	"strings"
	"time"
)

func subscribeProperty() {
	mqtt.Subscribe[map[string]any]("up/property/+", func(topic string, values *map[string]any) {
		topics := strings.Split(topic, "/")
		id := topics[2]

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

		dev.online = true
		dev.values["$online"] = true
		dev.values["$update"] = time.Now()
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
		dev.values[p.Name] = p.Value
	}
	dev.values["$online"] = true
	dev.values["$update"] = time.Now()

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
