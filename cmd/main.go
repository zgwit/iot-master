package main

import (
	"github.com/god-jason/bucket/config"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/service"
	"github.com/god-jason/bucket/web"
	master "github.com/zgwit/iot-master/v5"
	"github.com/zgwit/iot-master/v5/args"
)

// @title 物联大师接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @InfoInstanceName master
// @query.collection.format multi
func main() {
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
