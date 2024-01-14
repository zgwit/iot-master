package web

import (
	"github.com/zgwit/iot-master/v4/config"
)

const MODULE = "web"

func init() {
	config.Register(MODULE, "port", 8080)
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "cors", false)
	config.Register(MODULE, "gzip", true)
	config.Register(MODULE, "https", "")
	config.Register(MODULE, "cert", "")
	config.Register(MODULE, "key", "")
	config.Register(MODULE, "hosts", []string{})
	config.Register(MODULE, "email", "")

}
