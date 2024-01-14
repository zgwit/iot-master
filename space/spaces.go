package space

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/types"
)

var spaces lib.Map[Space]

func Ensure(id string) (*Space, error) {
	dev := spaces.Load(id)
	if dev == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		dev = spaces.Load(id)
	}
	return dev, nil
}

func Get(id string) *Space {
	return spaces.Load(id)
}

func Load(id string) error {
	var m types.Space
	has, err := db.Engine.ID(id).Get(&m)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到项目%s", id)
	}

	return From(&m)
}

func From(space *types.Space) error {
	p := New(space)

	spaces.Store(space.Id, p)
	//
	//err := db.Engine.Where("space_id = ?", id).And("disabled = ?", false).Find(&p.ExternalValidators)
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Engine.Where("space_id = ?", id).And("disabled = ?", false).Find(&p.ExternalAggregators)
	//if err != nil {
	//	return err
	//}

	var ds []types.SpaceDevice
	err := db.Engine.Where("space_id=?", space.Id).Find(&ds)
	if err != nil {
		return err
	}

	for _, d := range ds {
		dev := device.Get(d.DeviceId)
		if dev == nil {
			log.Error("设备未加载", d.DeviceId)
		} else {
			p.PutDevice(d.Name, dev)
		}
	}

	return nil
}

func Boot() error {
	//开机加载所有空间
	var ps []*types.Space
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = From(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
