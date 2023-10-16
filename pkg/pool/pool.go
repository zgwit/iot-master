package pool

import (
	"github.com/panjf2000/ants/v2"
	"github.com/zgwit/iot-master/v4/pkg/log"
)

var Pool *ants.Pool

func Open() (err error) {
	Pool, err = ants.NewPool(1000, ants.WithPanicHandler(func(err interface{}) {
		log.Error(err)
	}))
	return
}

func Close() {
	Pool.Release()
}

func Insert(task func()) error {
	return Pool.Submit(task)
}
