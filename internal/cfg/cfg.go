package cfg

import (
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/internal/broker"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/pool"
	"github.com/zgwit/iot-master/v4/web"
)

// Load 加载
func Load() {
	err := log.Load()
	if err != nil {
		_ = log.Store()
	}

	err = web.Load()
	if err != nil {
		_ = web.Store()
	}

	err = db.Load()
	if err != nil {
		_ = db.Store()
	}

	err = pool.Load()
	if err != nil {
		_ = pool.Store()
	}

	err = mqtt.Load()
	if err != nil {
		_ = mqtt.Store()
	}

	err = broker.Load()
	if err != nil {
		_ = broker.Store()
	}

}
