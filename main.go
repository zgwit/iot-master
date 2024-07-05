package main

import (
	_ "github.com/god-jason/bucket/action"
	_ "github.com/god-jason/bucket/aggregate"
	_ "github.com/god-jason/bucket/alarm"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/boot"
	_ "github.com/god-jason/bucket/device"
	_ "github.com/god-jason/bucket/function"
	_ "github.com/god-jason/bucket/gateway"
	_ "github.com/god-jason/bucket/history"
	"github.com/god-jason/bucket/pkg/service"
	_ "github.com/god-jason/bucket/product"
	_ "github.com/god-jason/bucket/scene"
	_ "github.com/god-jason/bucket/table"
	_ "github.com/god-jason/bucket/timer"
	_ "github.com/god-jason/bucket/user"
	"github.com/god-jason/bucket/web"
	"github.com/zgwit/iot-master/v5/args"
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
