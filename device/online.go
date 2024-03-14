package device

import (
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"strings"
)

func subscribeOnline() {

	mqtt.Subscribe("online/+", func(topic string, _ []byte) {
		topics := strings.Split(topic, "/")
		id := topics[1]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online()
	})

	mqtt.Subscribe("offline/+", func(topic string, _ []byte) {
		topics := strings.Split(topic, "/")
		id := topics[1]

		dev, err := Ensure(id)
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
		//	Product: dev.product.Assign,
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
