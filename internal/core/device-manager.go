package core

import (
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/internal/log"
	"github.com/zgwit/iot-master/model"
	"sync"
)

var allDevices sync.Map

func GetDevice(id uint64) *Device {
	d, ok := allDevices.Load(id)
	if ok {
		return d.(*Device)
	}
	return nil
}

func RemoveDevice(id uint64) error {
	d, ok := allDevices.LoadAndDelete(id)
	if ok {
		dev := d.(*Device)
		return dev.Stop()
	}
	return nil //error
}

func LoadDevices() error {
	return db.Store().ForEach(nil, func(d *model.Device) error {
		if d.Disabled {
			return nil
		}

		dev, err := NewDevice(d)
		if err != nil {
			log.Error(err)
			return nil
		}
		allDevices.Store(d.Id, dev)
		//err = dev.Start()
		//if err != nil {
		//	log.Error(err)
		//}
		return nil
	})
}

func LoadDevice(id uint64) (*Device, error) {
	var device model.Device

	err := db.Store().Get(id, &device)
	if err != nil {
		return nil, err
	}

	dev, err := NewDevice(&device)
	if err != nil {
		return dev, err
	}
	//allDevices.Store(device.Id, dev)
	allDevices.Store(id, dev)
	//err = dev.Start() //此处应该交给tunnel online启动
	return dev, nil
}
