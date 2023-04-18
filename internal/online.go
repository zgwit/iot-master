package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func subscribeOnline() error {
	mqtt.Client.Subscribe("online/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[2]
		id := topics[3]

		dev, err := GetDevice(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = true
		dev.Values["$online"] = true
	})

	mqtt.Client.Subscribe("offline/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[2]
		id := topics[3]

		dev, err := GetDevice(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = false
		dev.Values["$online"] = false

		//TODO 此处应该放置在 alarm中
		//alarm := types.Alarm{
		//	DeviceId: id,
		//	Type:     v.model.Type,
		//	Title:    v.model.Title,
		//	Level:    v.model.Level,
		//}
		//_, err = db.Engine.Insert(&alarm)
		//if err != nil {
		//	log.Error(err)
		//	//continue
		//}
		//
		//topic := fmt.Sprintf("alarm/%s/%s", pid, id)
		//payload, _ := json.Marshal(&alarm)
		//err = mqtt.Publish(topic, payload, false, 0)
	})

	return nil
}
