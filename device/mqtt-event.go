package device

import (
	"github.com/zgwit/iot-master/v4/payload"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"strings"
	"time"
)

func init() {
	db.Register(new(DeviceEvent))
}

type DeviceEvent struct {
	Id       int64          `json:"id"`
	DeviceId string         `json:"device_id" xorm:"index"`
	Name     string         `json:"name"`
	Label    string         `json:"label"`
	Output   map[string]any `json:"output" xorm:"json"`
	Created  time.Time      `json:"created" xorm:"created"`
}

func mqttEvent() {

	mqtt.SubscribeStruct[payload.Event]("device/+/event", func(topic string, event *payload.Event) {
		topics := strings.Split(topic, "/")
		//pid := topics[2]
		id := topics[1]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}

		//保存数据库
		_, _ = db.Engine.InsertOne(&DeviceEvent{
			DeviceId: id,
			Name:     event.Name,
			Label:    event.Title,
			Output:   event.Output,
		})

		switch event.Name {
		case "online":
			dev.Online = true
		case "offline":
			dev.Online = false
		}
	})

}
