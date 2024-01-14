package internal

import (
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"strings"
)

func subscribeOnline() {

	mqtt.Subscribe[any]("online/+", func(topic string, _ *any) {
		topics := strings.Split(topic, "/")
		id := topics[1]

		dev, err := device.Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online()
	})

	mqtt.Subscribe[any]("offline/+", func(topic string, _ *any) {
		topics := strings.Split(topic, "/")
		id := topics[1]

		dev, err := device.Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Offline()

		//产生日志
		//al := alarm.AlarmEx{
		//	Alarm: alarm.Alarm{
		//		DeviceId: id,
		//		Type:     "离线", //TODO 在 产品和设备 中配置
		//		Title:    "离线",
		//		Level:    3,
		//	},
		//	Product: dev.product.Name,
		//	Device:  dev.name,
		//}
		//_, err = db.Engine.Insert(&al.Alarm)
		//if err != nil {
		//	log.Error(err)
		//	//continue
		//}

		//通知
		//err = notify(&al)
		//if err != nil {
		//	log.Error(err)
		//	//continue
		//}

	})
}
