package main

import (
	"github.com/iot-master-contrib/alarm"
	"github.com/iot-master-contrib/classify"
	"github.com/iot-master-contrib/history"
	"github.com/zgwit/iot-master/v3"
	"github.com/zgwit/iot-master/v3/internal/app"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func main() {
	banner.Print()
	build.Println()

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

	err = classify.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	err = history.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	//注册静态页面
	fs := engine.FileSystem()

	master.Static(fs)
	alarm.Static(fs)
	classify.Static(fs)
	history.Static(fs)

	app.Register(alarm.App())
	app.Register(classify.App())
	app.Register(history.App())

	//启动
	engine.Serve()

}
