package cfg

import (
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/web"
)

// Load 加载
func Load() {
	_ = log.Load()
	_ = web.Load()
	_ = db.Load()
	_ = mqtt.Load()
	_ = broker.Load()
}
