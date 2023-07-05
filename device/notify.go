package device

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/payload"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

type sub struct {
	Id        string   `json:"id" xorm:"pk"`
	Name      string   `json:"name,omitempty"`
	Email     string   `json:"email,omitempty"`
	Cellphone string   `json:"cellphone,omitempty"`
	Channels  []string `json:"channels" xorm:"json"`
}

func notify(alarm *model.Alarm) error {
	//报警
	pa := payload.Alarm{
		Product: alarm.Product,
		Device:  alarm.Device,
		Type:    alarm.Type,
		Title:   alarm.Title,
		Level:   alarm.Level,
		Message: alarm.Message,
	}
	topic := fmt.Sprintf("alarm/%s/%s", alarm.ProductId, alarm.DeviceId)
	mqtt.Publish(topic, &pa)

	//找到订阅人
	var us []sub
	err := db.Engine.Table("subscription").
		Select("user.id, user.name, user.email, user.cellphone, subscription.channels").
		Join("INNER", "user", "user.id = subscription.user_id").
		Where("level<=?", alarm.Level).And("subscription.disabled!=1").
		And("product_id IS NULL OR product_id=\"\" OR product_id=?", alarm.ProductId).
		And("device_id IS NULL OR device_id=\"\" OR device_id=?", alarm.DeviceId).
		Find(&us)
	if err != nil {
		return err
	}

	//去除重复
	subs := map[string]sub{}
	for _, u := range us {
		if s, ok := subs[u.Id]; ok {
			for _, v := range u.Channels {
				found := false
				for _, vv := range s.Channels {
					if vv == v {
						found = true
					}
				}
				if !found {
					s.Channels = append(s.Channels, v)
				}
			}
		} else {
			subs[u.Id] = u
		}
	}

	//依次通知
	for _, u := range subs {
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
		//topic := fmt.Sprintf("notify/%s", u.Id)
		topic := fmt.Sprintf("notify/%s/%s", alarm.ProductId, alarm.DeviceId)
		mqtt.Publish(topic, &n)

		//不需要再广播了
		//nn := payload.Notify{
		//	Alarm:     pa,
		//	User:      u.Name,
		//	Email:     u.Email,
		//	Cellphone: u.Cellphone,
		//}
	}

	return nil
}
