package log

import (
	"github.com/zgwit/iot-master/v4/config"
)

const MODULE = "log"

func init() {
	config.Register(MODULE, "level", "trace")
	config.Register(MODULE, "caller", true)
	config.Register(MODULE, "text", true)
}
