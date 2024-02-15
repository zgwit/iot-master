package device

import (
	"github.com/zgwit/iot-master/v4/aggregator"
	"github.com/zgwit/iot-master/v4/alarm"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/event"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/types"
	"time"
)

type Device struct {
	id   string
	name string

	last   time.Time
	values map[string]any

	//事件监听
	EventData    event.Emitter[map[string]any]
	EventError   event.Emitter[error]
	EventAlarm   event.Emitter[*alarm.Alarm]
	EventOnline  event.Emitter[any]
	EventOffline event.Emitter[any]

	product *product.Product

	validators  []*alarm.Validator
	aggregators []aggregator.Aggregator
}

func (d *Device) Id() string {
	return d.id
}

func (d *Device) Name() string {
	return d.name
}

func (d *Device) Online() {
	d.values["$online"] = true
}

func (d *Device) Offline() {
	d.values["$online"] = false
}

func (d *Device) Values() map[string]any {
	return d.values
}

func (d *Device) createValidator(m *types.Validator) error {
	v, err := alarm.New(m)
	if err != nil {
		return err
	}
	d.validators = append(d.validators, v)
	return nil
}

//
//func (d *Device) Build() {
//	for _, v := range d.product.Validators {
//		err := d.createValidator(v)
//		if err != nil {
//			log.Error(err)
//		}
//	}
//	for _, v := range d.product.ExternalValidators {
//		err := d.createValidator(&v.Validator)
//		if err != nil {
//			log.Error(err)
//		}
//	}
//
//	var validators []*types.ExternalValidator
//	err := db.Engine.Where("device_id = ?", d.id).And("disabled = ?", false).Find(&validators)
//	if err != nil {
//		log.Error(err)
//	}
//	for _, v := range validators {
//		err := d.createValidator(&v.Validator)
//		if err != nil {
//			log.Error(err)
//		}
//	}
//
//}

func (d *Device) Push(values map[string]any) {
	for k, v := range values {
		d.values[k] = v
	}

	//数据聚合
	for _, a := range d.aggregators {
		err := a.Push(values)
		if err != nil {
			log.Error(err)
		}
	}

	//检查数据
	//d.Validate()
}

func (d *Device) Validate() {
	for _, v := range d.validators {
		ret := v.Validate(d.values)
		if !ret {
			//检查结果为真时，才产生报警
			continue
		}

		//入库
		al := alarm.AlarmEx{
			Alarm: alarm.Alarm{
				ProductId: d.product.Id,
				DeviceId:  d.id,
				Type:      v.Type,
				Title:     v.Title,
				Level:     v.Level,
				Message:   v.Template, //TODO 模板格式化
			},
			Product: d.product.Name,
			Device:  d.name,
		}
		_, err := db.Engine.Insert(&al.Alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

		//通知
		//err = internal.notify(&al)
		//if err != nil {
		//	log.Error(err)
		//	//continue
		//}
	}
}

func New(m *types.Device) *Device {
	//time.Now().Unix()
	return &Device{
		id:     m.Id,
		name:   m.Name,
		values: make(map[string]any),
	}
}
