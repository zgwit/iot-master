package device

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/aggregator"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/internal/product"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/model"
	"github.com/zgwit/iot-master/v4/validator"
	"time"
)

var devices lib.Map[Device]

type Device struct {
	*model.Device

	Last   time.Time
	Values map[string]any

	product *product.Product

	validators  []*validator.Validator
	aggregators []aggregator.Aggregator
}

func (d *Device) createValidator(m *model.ModValidator) error {
	v, err := validator.New(m)
	if err != nil {
		return err
	}
	d.validators = append(d.validators, v)
	return nil
}

func (d *Device) createAggregator(m *model.ModAggregator) error {
	a, err := aggregator.New(m, func(val float64, err error) {
		if err != nil {
			log.Error(err)
			return
		}
		his := model.History{
			DeviceId: d.Id,
			Point:    m.Assign,
			Value:    val,
			Time:     time.Now(),
		}
		_, err = db.Engine.InsertOne(&his)
		if err != nil {
			log.Error(err)
		}
	})
	if err != nil {
		return err
	}
	d.aggregators = append(d.aggregators, a)
	return nil
}

func (d *Device) Build() {
	for _, v := range d.product.Validators {
		err := d.createValidator(v)
		if err != nil {
			log.Error(err)
		}
	}
	for _, v := range d.product.ExternalValidators {
		err := d.createValidator(&v.ModValidator)
		if err != nil {
			log.Error(err)
		}
	}

	var validators []*model.Validator
	err := db.Engine.Where("device_id = ?", d.Id).And("disabled = ?", false).Find(&validators)
	if err != nil {
		log.Error(err)
	}
	for _, v := range validators {
		err := d.createValidator(&v.ModValidator)
		if err != nil {
			log.Error(err)
		}
	}

	for _, v := range d.product.Aggregators {
		err := d.createAggregator(v)
		if err != nil {
			log.Error(err)
		}
	}
	for _, v := range d.product.ExternalAggregators {
		err := d.createAggregator(&v.ModAggregator)
		if err != nil {
			log.Error(err)
		}
	}

	var aggregators []*model.Aggregator
	err = db.Engine.Where("device_id = ?", d.Id).And("disabled = ?", false).Find(&aggregators)
	if err != nil {
		log.Error(err)
	}
	for _, v := range aggregators {
		err := d.createAggregator(&v.ModAggregator)
		if err != nil {
			log.Error(err)
		}
	}

}

func (d *Device) Push(values map[string]any) {
	for k, v := range values {
		d.Values[k] = v
	}

	//数据聚合
	for _, a := range d.aggregators {
		err := a.Push(values)
		if err != nil {
			log.Error(err)
		}
	}

	//检查数据
	d.Validate()
}

func (d *Device) Validate() {
	for _, v := range d.validators {
		ret := v.Validate(d.Values)
		if !ret {
			//检查结果为真时，才产生报警
			continue
		}

		//入库
		alarm := model.AlarmEx{
			Alarm: model.Alarm{
				ProductId: d.product.Id,
				DeviceId:  d.Id,
				Type:      v.Type,
				Title:     v.Title,
				Level:     v.Level,
				Message:   v.Template, //TODO 模板格式化
			},
			Product: d.product.Name,
			Device:  d.Name,
		}
		_, err := db.Engine.Insert(&alarm.Alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

		//通知
		err = notify(&alarm)
		if err != nil {
			log.Error(err)
			//continue
		}
	}
}

func New(m *model.Device) *Device {
	//time.Now().Unix()
	return &Device{
		Device: m,
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
	d := New(device)

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

	//构建
	d.Build()

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
