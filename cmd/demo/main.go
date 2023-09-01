package main

import (
	"github.com/iot-master-contrib/aliyun"
	"github.com/iot-master-contrib/classify"
	"github.com/iot-master-contrib/ipc"
	"github.com/iot-master-contrib/modbus"
	"github.com/iot-master-contrib/scada"
	"github.com/iot-master-contrib/tsdb"
	"github.com/zgwit/iot-master/v4/app"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/pkg/banner"
	"github.com/zgwit/iot-master/v4/pkg/build"
	"github.com/zgwit/iot-master/v4/web"
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

	err = classify.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	err = tsdb.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	err = ipc.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	err = modbus.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	scada.Route(engine)
	err = scada.Sync()
	if err != nil {
		log.Fatal(err)
	}

	err = aliyun.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}
	_ = db.Engine.Sync2(aliyun.Models()...)

	//注册静态页面
	fs := engine.FileSystem()

	master.Static(fs)
	classify.Static(fs)
	tsdb.Static(fs)
	ipc.Static(fs)
	modbus.Static(fs)
	scada.Static(fs)
	aliyun.Static(fs)

	app.Register(classify.App())
	app.Register(tsdb.App())
	app.Register(ipc.App())
	app.Register(modbus.App())
	app.Register(scada.App())
	app.Register(aliyun.App())

	//启动
	engine.Serve()

}
