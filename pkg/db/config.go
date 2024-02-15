package db

import (
	"github.com/zgwit/iot-master/v4/pkg/config"
)

const MODULE = "database"

func init() {
	config.Register(MODULE, "type", "mysql")
	config.Register(MODULE, "url", "root:root@tcp(localhost:3306)/master?charset=utf8")
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "sync", true)
}
