package device

import (
	"encoding/json"
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/payload"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func SubscribeOnline() error {
	mqtt.Client.Subscribe("online/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[1]
		id := topics[2]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = true
		dev.Values["$online"] = true
	})

	mqtt.Client.Subscribe("offline/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		pid := topics[1]
		id := topics[2]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = false
		dev.Values["$online"] = false

		//产生日志
		alarm := model.Alarm{
			DeviceId: id, //TODO 配置化
			Type:     "离线",
			Title:    "离线",
			Level:    3,
		}
		_, err = db.Engine.Insert(&alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

		//广播报警内容
		topic := fmt.Sprintf("alarm/%s/%s", pid, id)
		data, _ := json.Marshal(&payload.Alarm{
			Product: dev.product.Name,
			Device:  dev.Name,
			Type:    alarm.Type,
			Title:   alarm.Title,
			Message: alarm.Message,
			Level:   alarm.Level,
		})
		err = mqtt.Publish(topic, data, false, 0)

	})

	return nil
}
