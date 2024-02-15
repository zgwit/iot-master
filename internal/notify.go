package internal

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/alarm"
	"github.com/zgwit/iot-master/v4/payload"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
)

type sub struct {
	Id        string   `json:"id" xorm:"pk"`
	Name      string   `json:"name,omitempty"`
	Email     string   `json:"email,omitempty"`
	Cellphone string   `json:"cellphone,omitempty"`
	Channels  []string `json:"channels" xorm:"json"`
}

func notify(al *alarm.AlarmEx) error {
	//报警
	pa := payload.Alarm{
		Product: al.Product,
		Device:  al.Device,
		Type:    al.Type,
		Title:   al.Title,
		Level:   al.Level,
		Message: al.Message,
	}
	topic := fmt.Sprintf("al/%s/%s", al.ProductId, al.DeviceId)
	mqtt.Publish(topic, &pa)

	//找到订阅人
	var us []sub
	err := db.Engine.Table("subscription").
		Select("user.id, user.name, user.email, user.cellphone, subscription.channels").
		Join("INNER", "user", "user.id = subscription.user_id").
		Where("level<=?", al.Level).And("subscription.disabled!=1").
		And("product_id IS NULL OR product_id=\"\" OR product_id=?", al.ProductId).
		And("device_id IS NULL OR device_id=\"\" OR device_id=?", al.DeviceId).
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
		n := alarm.Notification{
			AlarmId:  al.Id,
			UserId:   u.Id,
			Channels: u.Channels,
		}

		//保存记录
		_, err = db.Engine.InsertOne(&n)
		if err != nil {
			return err
		}

		//MQTT通知，第三方插件来发送
		//topic := fmt.Sprintf("notify/%s", u.id)
		topic := fmt.Sprintf("notify/%s/%s", al.ProductId, al.DeviceId)
		mqtt.Publish(topic, &n)

		//不需要再广播了
		//nn := payload.Notify{
		//	Alarm:     pa,
		//	User:      u.name,
		//	Email:     u.Email,
		//	Cellphone: u.Cellphone,
		//}
	}

	return nil
}
