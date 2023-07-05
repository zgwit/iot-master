package device

import (
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func SubscribeOnline() error {
	mqtt.Subscribe("online/+/+", func(topic string, payload []byte) {
		topics := strings.Split(topic, "/")
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

	mqtt.Subscribe("offline/+/+", func(topic string, payload []byte) {
		topics := strings.Split(topic, "/")
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
			ProductId: pid,
			DeviceId:  id,
			Product:   dev.product.Name,
			Device:    dev.Name,
			Type:      "离线", //TODO 在 产品和设备 中配置
			Title:     "离线",
			Level:     3,
		}
		_, err = db.Engine.Insert(&alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

		//通知
		err = notify(&alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

	})

	return nil
}
