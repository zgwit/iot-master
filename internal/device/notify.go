package device

import (
	"encoding/json"
	"fmt"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

type sub struct {
	Id        string   `json:"id" xorm:"pk"`
	Email     string   `json:"email,omitempty"`
	Cellphone string   `json:"cellphone,omitempty"`
	Channels  []string `json:"channels" xorm:"json"`
}

func notify(alarm *model.Alarm) error {
	//alarm.

	//找到订阅人
	var us []sub
	err := db.Engine.Table("subscription").
		Select("user.id, user.email, user.cellphone, subscription.channels").
		Join("INNER", "user", "user.id = subscription.user_id").
		Where("level<=?", alarm.Level).And("disabled<>1").
		And("( product_id IS NULL OR product_id=`` OR product_id=?)", alarm.ProductId).
		And("( device_id IS NULL OR device_id=`` OR device_id=?)", alarm.DeviceId).
		Find(&us)
	if err != nil {
		return err
	}

	//TODO 去除重复的？？？

	//依次通知
	for _, u := range us {
		n := model.Notification{
			AlarmId:  alarm.Id,
			UserId:   u.Id,
			Channels: u.Channels,
		}

		//保存记录
		_, err = db.Engine.InsertOne(&n)
		if err != nil {
			return err
		}

		//MQTT通知，第三方插件来发送
		topic := fmt.Sprintf("notify/%d", n.Id)
		data, _ := json.Marshal(&n)
		err = mqtt.Publish(topic, data, false, 0)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
