package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func subscribeEvent() error {
	mqtt.Client.Subscribe("up/event/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		//pid := topics[2]
		id := topics[3]

		dev, err := GetDevice(id)
		if err != nil {
			log.Error(err)
			return
		}

		var payload model.PayloadEvent
		err = json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			return
		}

		//保存数据库
		_, _ = db.Engine.InsertOne(model.DeviceEvent{
			DeviceId: id,
			Name:     payload.Name,
			Label:    payload.Title,
			Output:   payload.Output,
		})

		switch payload.Name {
		case "online":
			dev.Values["$online"] = true
		case "offline":
			dev.Values["$online"] = false
		}
	})

	return nil
}
