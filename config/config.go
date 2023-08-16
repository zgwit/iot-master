package config

import (
	"github.com/zgwit/iot-master/v3/broker"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/oem"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

// Load 加载
func Load() {
	_ = log.Load()
	_ = web.Load()
	_ = oem.Load()
	_ = db.Load()
	_ = mqtt.Load()
	_ = broker.Load()
}
