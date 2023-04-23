package config

import (
	"github.com/zgwit/iot-master/v3/internal/broker"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/oem"
	"github.com/zgwit/iot-master/v3/pkg/web"
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

	err = oem.Load()
	if err != nil {
		_ = oem.Store()
	}

	err = db.Load()
	if err != nil {
		_ = db.Store()
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
