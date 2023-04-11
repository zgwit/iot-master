package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"time"
)

var Devices lib.Map[Device]

type Device struct {
	Id          string
	Online      bool
	Last        time.Time
	Values      map[string]any
	product     *Product
	constraints []*constraint
}

type constraint struct {
	model *model.ModConstraint
	eval  gval.Evaluable //当修改产品信息时，需要同步设备参数，用 重载？？？
	//again      bool
	start int64 //开始时间s
	total uint  //报警次数
}

func NewDevice(id string) *Device {
	//time.Now().Unix()
	return &Device{
		Id:          id,
		Values:      make(map[string]any),
		constraints: make([]*constraint, 0),
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
	}

	p := Products.Load(device.ProductId)
	if p == nil {
		return fmt.Errorf("product %s not found", device.ProductId)
	}

	//复制基础变量
	for k, v := range p.values {
		d.Values[k] = v
	}
	//复制设备变量
	for k, v := range device.Parameters {
		d.Values[k] = v
	}

	//构建约束器
	for k, v := range p.eval {
		c := &constraint{
			model: &p.model.Constraints[k],
			eval:  v,
		}
		d.constraints = append(d.constraints, c)
	}

	Devices.Store(device.Id, d)
	return nil
}

func (d *Device) Constrain() {
	now := time.Now().Unix()

	for _, e := range d.constraints {
		ret, err := e.eval.EvalBool(context.Background(), d.Values)
		if err != nil {
			log.Error(err)
			continue
			//return
		}
		if ret {
			//约束OK，检查下一个
			e.total = 0
			e.start = 0
			continue
		}

		cs := e.model

		//now := time.Now().Unix()
		if e.start == 0 {
			e.start = now
		}

		//延迟报警
		if cs.Delay > 0 {
			if now < e.start+int64(cs.Delay) {
				continue
			}
		}

		//重复报警
		if cs.Again > 0 && e.total < cs.Total {
			if now < e.start+int64(cs.Again) {
				continue
			}

			//重置开始时间
			e.start = now // + int64(cs.Delay)
			e.total++
		}

		//报警
		alarm := &model.Alarm{
			DeviceId: d.Id,
			Level:    cs.Level,
			Title:    cs.Title,
			Message:  "",
		}

		//入库
		_, err = db.Engine.InsertOne(alarm)
		if err != nil {
			log.Error(err)
		}

		//mqtt广播
		topic := fmt.Sprintf("alarm/%s/%s", d.product.model.Id, d.Id)
		payload, _ := json.Marshal(alarm)
		err = mqtt.Publish(topic, payload, false, 0)
		if err != nil {
			log.Error(err)
		}
	}
}
