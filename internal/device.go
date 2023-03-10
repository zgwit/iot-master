package internal

import (
	"context"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"time"
)

var Devices lib.Map[Device]

type Device struct {
	Id         string
	Online     bool
	Last       time.Time
	Properties map[string]any
	Product    *Product
	checkers   map[int]*checker
}

type checker struct {
	again bool
	start int64 //

}

func NewDevice(id string) *Device {
	//time.Now().Unix()
	return &Device{
		Id:         id,
		Properties: make(map[string]any),
		checkers:   make(map[int]*checker),
	}
}

func (d *Device) Constrain() {
	for i, e := range d.Product.eval {
		ret, err := e.EvalBool(context.Background(), d.Properties)
		if err != nil {
			log.Error(err)
			continue
			//return
		}
		if ret {
			//约束OK，检查下一个
			continue
		}

		c, ok := d.checkers[i]
		if !ok {
			c = &checker{}
			d.checkers[i] = c
		}

		cs := d.Product.model.Constraints[i]

		now := time.Now().Unix()
		//延迟报警
		if cs.Delay > 0 {
			if c.start+int64(cs.Delay) > now {
				continue
			}
		}

		//再次报警
		if cs.Again > 0 {
			if c.start+int64(cs.Again) > now {
				c.start = now + int64(cs.Delay)
				continue
			}
		}

	}
}
