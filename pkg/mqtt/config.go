package mqtt

import (
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/config"
)

const MODULE = "mqtt"

func init() {
	config.Register(MODULE, "url", "mqtt://localhost:1843")
	config.Register(MODULE, "clientId", lib.RandomString(8))
	config.Register(MODULE, "username", "")
	config.Register(MODULE, "password", "")
}
