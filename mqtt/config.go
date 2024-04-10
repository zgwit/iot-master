package mqtt

import (
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/lib"
)

const MODULE = "mqtt"

func init() {
	config.Register(MODULE, "url", "mqtt://localhost:1843")
	config.Register(MODULE, "clientId", lib.RandomString(8))
	config.Register(MODULE, "username", "")
	config.Register(MODULE, "password", "")
}
