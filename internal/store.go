package internal

import (
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/convert"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

func store() {
	var stores []model.History
	Devices.Range(func(id string, dev *Device) bool {
		//如果没在线，就不保存
		if !dev.Online {
			return true
		}

		//检查需要保存的变量
		for k, s := range dev.product.stores {
			if v, ok := dev.Values[k]; ok {
				value := convert.ToFloat64(v)
				if s == "save" {
					stores = append(stores, model.History{
						DeviceId: dev.Id,
						Point:    k,
						Value:    value,
						Time:     model.Time{},
					})
				} else if s == "diff" {
					if l, ok := dev.last[k]; ok {
						stores = append(stores, model.History{
							DeviceId: dev.Id,
							Point:    k,
							Value:    value - l,
							Time:     model.Time{},
						})

					}
					dev.last[k] = value
				}
			}

		}
		return true
	})

	n, err := db.Engine.Insert(stores)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("store ", n)
	}
}
