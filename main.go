package main

import (
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
)

func main() {

	_ = database.Open(&config.Config.Database)

	_ = tsdb.Open(&config.Config.History)

	_ = connect.LoadTunnels()

	_ = master.Start()

	//TODO，判断是否开启Web
	web.Serve(&config.Config.Web)
}
