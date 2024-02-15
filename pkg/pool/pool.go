package pool

import (
	"github.com/panjf2000/ants/v2"
	"github.com/zgwit/iot-master/v4/pkg/config"
	"github.com/zgwit/iot-master/v4/pkg/log"
)

var Pool *ants.Pool

func Open() (err error) {
	Pool, err = ants.NewPool(config.GetInt(MODULE, "size"), ants.WithPanicHandler(func(err interface{}) {
		log.Error(err)
	}))
	return
}

func Close() {
	Pool.Release()
	Pool = nil
}

func Insert(task func()) error {
	return Pool.Submit(task)
}
