package notify

import (
	"github.com/zgwit/iot-master/v4/alarm"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/pool"
)

type Notifier interface {
	Notify(notification *Notification) error
}

var notifiers = map[string]Notifier{}

func Register(channel string, notifier Notifier) {
	//notifiers = append(notifiers, h)
	notifiers[channel] = notifier
}

func Notify(alarm *alarm.Alarm) error {

	//找到订阅人
	var subs []Subscription
	err := db.Engine.Where("level<=?", alarm.Level).And("disabled!=1").
		And("product_id IS NULL OR product_id=\"\" OR product_id=?", alarm.ProductId).
		And("project_id IS NULL OR product_id=\"\" OR project_id=?", alarm.ProjectId).
		And("space_id IS NULL OR product_id=\"\" OR space_id=?", alarm.SpaceId).
		And("device_id IS NULL OR device_id=\"\" OR device_id=?", alarm.DeviceId).
		Find(&subs)
	if err != nil {
		return err
	}

	//创建通知 并 去除重复
	ns := map[string]*Notification{}
	for _, sub := range subs {
		if n, ok := ns[sub.UserId]; ok {
			for _, c := range sub.Channels {
				found := false
				for _, cc := range n.Channels {
					if cc == c {
						found = true
					}
				}
				if !found {
					n.Channels = append(n.Channels, c)
				}
			}
		} else {
			ns[sub.UserId] = &Notification{
				AlarmId:  alarm.Id,
				Title:    alarm.Title,
				UserId:   sub.UserId,
				User:     sub.User,
				Channels: sub.Channels,
			}
		}
	}

	//进行通知
	for _, n := range ns {
		_, err := db.Engine.InsertOne(n)
		if err != nil {
			log.Error(err)
		}

		//每个通道逐一通讯
		for _, c := range n.Channels {
			if notifier, ok := notifiers[c]; ok {
				err := pool.Insert(func() {
					err := notifier.Notify(n)
					if err != nil {
						log.Error(err)
					}
				})
				if err != nil {
					log.Error(err)
				}
			}
		}
	}

	return nil
}
