package core

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"time"
)

var Devices lib.Map[Device]

type Device struct {
	Id      string
	Online  bool
	Last    time.Time
	Values  map[string]any
	last    map[string]float64
	product *Product
}

func NewDevice(id string) *Device {
	//time.Now().Unix()
	return &Device{
		Id:     id,
		Values: make(map[string]any),
		last:   make(map[string]float64),
	}
}

func GetDevice(id string) (*Device, error) {
	dev := Devices.Load(id)
	if dev == nil {
		//log.Infof("加载设备 %s", id)
		//加载设备
		err := LoadDeviceById(id)
		if err != nil {
			return nil, err
		}
		dev = Devices.Load(id)
	}
	return dev, nil
}

func LoadDeviceById(id string) error {
	var dev model.Device
	get, err := db.Engine.ID(id).Get(&dev)
	if err != nil {
		return err
	}
	if !get {
		return fmt.Errorf("device %s not found", id)
	}
	return LoadDevice(&dev)
}

func LoadDevice(device *model.Device) error {
	d := &Device{
		Id:     device.Id,
		Values: make(map[string]any),
		last:   make(map[string]float64),
	}

	//绑定产品
	p := Products.Load(device.ProductId)
	if p == nil {
		return fmt.Errorf("product %s not found", device.ProductId)
	}
	d.product = p

	//复制基础变量
	for k, v := range p.values {
		d.Values[k] = v
	}
	//复制设备变量
	for k, v := range device.Parameters {
		d.Values[k] = v
	}

	Devices.Store(device.Id, d)
	return nil
}
