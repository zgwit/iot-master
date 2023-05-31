package main

import (
	"github.com/iot-master-contrib/alarm"
	"github.com/iot-master-contrib/classify"
	"github.com/iot-master-contrib/history"
	"github.com/iot-master-contrib/influxdb"
	"github.com/iot-master-contrib/ipc"
	"github.com/iot-master-contrib/modbus"
	"github.com/iot-master-contrib/scada"
	"github.com/iot-master-contrib/sms"
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

	err = influxdb.Startup(engine)
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

	err = scada.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	err = sms.Startup(engine)
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
	influxdb.Static(fs)
	ipc.Static(fs)
	modbus.Static(fs)
	scada.Static(fs)
	sms.Static(fs)

	app.Register(alarm.App())
	app.Register(classify.App())
	app.Register(history.App())
	app.Register(influxdb.App())
	app.Register(ipc.App())
	app.Register(modbus.App())
	app.Register(scada.App())
	app.Register(sms.App())

	//启动
	engine.Serve()

}
