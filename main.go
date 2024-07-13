package main

import (
	_ "github.com/god-jason/bucket/aggregate"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/boot"
	_ "github.com/god-jason/bucket/function"
	"github.com/god-jason/bucket/pkg/service"
	_ "github.com/god-jason/bucket/table"
	"github.com/god-jason/bucket/web"
	_ "github.com/zgwit/iot-master/v5/action"
	_ "github.com/zgwit/iot-master/v5/alarm"
	"github.com/zgwit/iot-master/v5/args"
	_ "github.com/zgwit/iot-master/v5/device"
	_ "github.com/zgwit/iot-master/v5/gateway"
	_ "github.com/zgwit/iot-master/v5/history"
	_ "github.com/zgwit/iot-master/v5/product"
	_ "github.com/zgwit/iot-master/v5/scene"
	_ "github.com/zgwit/iot-master/v5/timer"
	_ "github.com/zgwit/iot-master/v5/user"
	"log"
)

func main() {
	args.Parse()

	err := service.Register(Startup, Shutdown)
	if err != nil {
		log.Fatal(err)
	}

	if args.Install {
		err = service.Install()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if args.Uninstall {
		err = service.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = service.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Startup() error {
	err := boot.Startup()
	if err != nil {
		return err
	}

	//注册前端接口
	api.RegisterRoutes(web.Engine.Group("/api"))

	//监听Websocket
	//web.Engine.GET("/mqtt", broker.GinBridge)

	//附件
	//web.Engine.Static("/static", "static")
	//web.Engine.Static("/attach", filepath.Join(viper.GetString("data"), "attach"))

	//前端 移入子工程 github.com/iot-master-contrib/webui
	//web.Static.Put("", http.FS(wwwFiles), "www", "index.html")

	//注册静态文件
	web.Static.PutDir("", "www", "", "index.html")

	return web.Serve()
}

func Shutdown() error {
	return boot.Shutdown()
}
