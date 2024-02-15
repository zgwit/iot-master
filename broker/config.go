package broker

import (
	"github.com/zgwit/iot-master/v4/pkg/config"
)

const MODULE = "broker"

func init() {
	config.Register(MODULE, "enable", true)
	config.Register(MODULE, "port", 1843)
}
