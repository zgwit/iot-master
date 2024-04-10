package pool

import (
	"github.com/zgwit/iot-master/v4/config"
)

const MODULE = "pool"

func init() {
	config.Register(MODULE, "size", 10000)
}
