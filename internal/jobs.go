package internal

import (
	"github.com/robfig/cron/v3"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/convert"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"time"
)

var _cron *cron.Cron

func StartJobs() error {
	_cron = cron.New()
	_, err := _cron.AddFunc("* * * * *", store) //测试，每分钟执行一次
	//_, err := _cron.AddFunc("@hourly", store)
	if err != nil {
		return err
	}

	_cron.Start()
	return nil
}

func StopJobs() {
	_cron.Stop()
}

func store() {
	log.Info("store job")

	var stores []model.History
	Devices.Range(func(id string, dev *Device) bool {
		//如果没在线，就不保存
		if !dev.Online {
			return true
		}

		now := model.Time(time.Now())

		//检查需要保存的变量
		for k, s := range dev.product.stores {
			if v, ok := dev.Values[k]; ok {
				value := convert.ToFloat64(v)
				if s == "save" {
					stores = append(stores, model.History{
						DeviceId: dev.Id,
						Point:    k,
						Value:    value,
						Time:     now,
					})
				} else if s == "increase" {
					if l, ok := dev.last[k]; ok {
						stores = append(stores, model.History{
							DeviceId: dev.Id,
							Point:    k,
							Value:    value - l,
							Time:     now,
						})
					}
					dev.last[k] = value
				}
			}

		}
		return true
	})

	if len(stores) > 0 {
		n, err := db.Engine.Insert(stores)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("store ", n)
		}
	}
}
