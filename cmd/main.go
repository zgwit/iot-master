package main

import (
	_ "github.com/iot-master-contrib/camera"
	_ "github.com/iot-master-contrib/influxdb"
	_ "github.com/iot-master-contrib/nuwa"
	_ "github.com/iot-master-contrib/webui"
	master "github.com/zgwit/iot-master/v4"
	"github.com/zgwit/iot-master/v4/args"
	"github.com/zgwit/iot-master/v4/config"
	_ "github.com/zgwit/iot-master/v4/docs"
	"github.com/zgwit/iot-master/v4/log"
	_ "github.com/zgwit/iot-master/v4/modbus"
	"github.com/zgwit/iot-master/v4/pkg/build"
	"github.com/zgwit/iot-master/v4/pkg/service"
	"github.com/zgwit/iot-master/v4/web"
)

// @title 物联大师接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @InfoInstanceName master
// @query.collection.format multi
func main() {
	build.Println()

	args.Parse()

	config.Name("iot-master")

	//传递参数到服务
	//serviceConfig.Arguments = []string{"-c", args.ConfigPath}

	err := service.Register(func() {

		err := master.Startup()
		if err != nil {
			log.Error(err)
			return
		}

		_ = web.Serve()
	}, func() {
		err := master.Shutdown()
		if err != nil {
			log.Error(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	if args.Uninstall {
		err = service.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("卸载服务成功")
		return
	}

	if args.Install {
		err = service.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("安装服务成功")
		return
	}

	err = service.Run()
	if err != nil {
		log.Error(err)
	}
}
