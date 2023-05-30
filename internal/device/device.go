package device

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/internal/product"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"time"
)

var devices lib.Map[Device]

type Device struct {
	Id      string
	Online  bool
	Last    time.Time
	Values  map[string]any
	product *product.Product
}

func New(id string) *Device {
	//time.Now().Unix()
	return &Device{
		Id:     id,
		Values: make(map[string]any),
	}
}

func Ensure(id string) (*Device, error) {
	dev := devices.Load(id)
	if dev == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		dev = devices.Load(id)
	}
	return dev, nil
}

func Get(id string) *Device {
	return devices.Load(id)
}

func Load(id string) error {
	var dev model.Device
	get, err := db.Engine.ID(id).Get(&dev)
	if err != nil {
		return err
	}
	if !get {
		return fmt.Errorf("device %s not found", id)
	}
	return From(&dev)
}

func From(device *model.Device) error {
	d := New(device.Id)

	//绑定产品
	p, err := product.Ensure(device.ProductId)
	if err != nil {
		return err
	}
	d.product = p

	//复制基础参数
	for _, v := range p.Parameters {
		d.Values[v.Name] = v.Default
	}

	//复制设备参数
	for k, v := range device.Parameters {
		d.Values[k] = v
	}

	devices.Store(device.Id, d)
	return nil
}

func GetOnlineCount() int64 {
	var count int64 = 0
	devices.Range(func(_ string, dev *Device) bool {
		count++
		return true
	})
	return count
}
