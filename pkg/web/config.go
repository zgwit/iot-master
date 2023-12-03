package web

import (
	"github.com/zgwit/iot-master/v4/config"
)

const MODULE = "web"

func init() {
	config.Register(MODULE, "addr", ":8080")
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "cors", false)
	config.Register(MODULE, "gzip", true)
}
