package space

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
)

var spaces lib.Map[Space]

func Get(id string) *Space {
	return spaces.Load(id)
}

func Load(id string) error {
	var m Space
	has, err := db.Engine.ID(id).Get(&m)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到项目%s", id)
	}
	return From(&m)
}

func From(m *Space) error {
	spaces.Store(m.Id, m)

	var ds []*SpaceDevice
	err := db.Engine.Where("space_id=?", m.Id).Find(&ds)
	if err != nil {
		return err
	}

	for _, d := range ds {
		dev := device.Get(d.DeviceId)
		if dev == nil {
			log.Error("设备未加载", d.DeviceId)
		} else {
			m.PutDevice(d.Name, dev)
		}
	}
	//TODO 设备上线，自动找空间

	return nil
}
