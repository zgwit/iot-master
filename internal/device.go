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
	Id         string
	Online     bool
	Last       time.Time
	Properties map[string]any
	Product    *Product
	checkers   []*Constraint
}

type Constraint struct {
	constraint *model.ModConstraint
	eval       gval.Evaluable //当修改产品信息时，需要同步设备参数，用 重载？？？
	again      bool
	start      int64 //
	total      uint
}

func NewDevice(id string) *Device {
	//time.Now().Unix()
	return &Device{
		Id:         id,
		Properties: make(map[string]any),
		checkers:   make([]*Constraint, 0),
	}
}

func (d *Device) Constrain() {
	for _, e := range d.checkers {
		ret, err := e.eval.EvalBool(context.Background(), d.Properties)
		if err != nil {
			log.Error(err)
			continue
			//return
		}
		if ret {
			//约束OK，检查下一个
			e.total = 0
			continue
		}

		cs := e.constraint

		now := time.Now().Unix()
		//延迟报警
		if cs.Delay > 0 {
			if e.start+int64(cs.Delay) > now {
				continue
			}
		}

		//重复报警
		if cs.Again > 0 && e.total < cs.Total {
			if e.start+int64(cs.Again) > now {
				e.start = now + int64(cs.Delay)
				continue
			}
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
		topic := fmt.Sprintf("alarm/%s/%s", d.Product.model.Id, d.Id)
		payload, _ := json.Marshal(alarm)
		err = mqtt.Publish(topic, payload, false, 0)
		if err != nil {
			log.Error(err)
		}
	}
}
