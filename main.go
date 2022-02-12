package main

import (
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/internal"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
)

func main() {

	_ = internal.Start()

	_ = database.Open(&config.Config.Database)

	_ = tsdb.Open(&config.Config.History)

	web.Serve(&config.Config.Web)
}
