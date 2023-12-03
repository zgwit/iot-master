package broker

import (
	"github.com/zgwit/iot-master/v4/config"
)

const MODULE = "broker"

func init() {
	config.Register(MODULE, "enable", "true")
	config.Register(MODULE, "addr", ":1843")
}
