package pool

import "github.com/panjf2000/ants/v2"

var Pool *ants.Pool

func Open() (err error) {
	Pool, err = ants.NewPool(1000)
	return
}

func Close() {
	Pool.Release()
}

func Insert(task func()) error {
	return Pool.Submit(task)
}
