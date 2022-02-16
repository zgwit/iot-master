package main

import (
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
	"log"
	"time"
)

func main() {

	err := database.Open(config.Config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	user := model.User{
		Username: "admin",
		Nickname: "admin",
		Created:  time.Now(),
	}

	err = database.Master.Save(&user)
	if err != nil {
		log.Fatal(err)
	}

	password := model.Password{
		Id:       user.Id,
		Password: "123456",
	}

	err = database.Master.Save(&password)
	if err != nil {
		log.Fatal(err)
	}

	err = tsdb.Open(config.Config.History)
	if err != nil {
		log.Fatal(err)
	}
	defer tsdb.Close()

	err = connect.LoadTunnels()
	if err != nil {
		log.Fatal(err)
	}
	//defer connect.Close()

	err = master.Start()
	if err != nil {
		log.Fatal(err)
	}
	//defer master.Close()

	//TODO，判断是否开启Web
	web.Serve(config.Config.Web)
}
