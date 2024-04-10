package pool

import (
	"github.com/panjf2000/ants/v2"
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/log"
)

var Pool *ants.Pool

func Startup() (err error) {
	Pool, err = ants.NewPool(config.GetInt(MODULE, "size"), ants.WithPanicHandler(func(err interface{}) {
		log.Error(err)
	}))
	return
}

func Shutdown() error {
	if Pool != nil {
		Pool.Release()
		Pool = nil
	}
	return nil
}

func Insert(task func()) error {
	if Pool == nil {
		go task()
	}
	return Pool.Submit(task)
}
