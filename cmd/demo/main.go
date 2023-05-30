package main

import (
	"github.com/iot-master-contrib/alarm"
	"github.com/zgwit/iot-master/v3"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func main() {

	//原本的Main函数
	engine := web.CreateEngine()

	//启动
	err := master.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	err = alarm.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	//注册静态页面
	fs := engine.FileSystem()

	master.Static(fs)
	alarm.Static(fs)

	//启动
	engine.Serve()

}
