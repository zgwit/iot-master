package main

import (
	"embed"
	"github.com/kardianos/service"
	_ "github.com/zgwit/iot-master/v3/docs"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/internal/api"
	"github.com/zgwit/iot-master/v3/internal/app"
	"github.com/zgwit/iot-master/v3/internal/args"
	"github.com/zgwit/iot-master/v3/internal/broker"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:embed all:www
var wwwFiles embed.FS

var serviceConfig = &service.Config{
	Name:        "iot-master",
	DisplayName: "物联大师",
	Description: "物联网数据中台",
	Arguments:   nil,
}

// @title 物联大师接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @query.collection.format multi
func main() {
	args.Parse()

	//传递参数到服务
	serviceConfig.Arguments = []string{"-c", args.ConfigPath}

	// 构建服务对象
	program := &Program{}
	s, err := service.New(program, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if args.Uninstall {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("卸载服务成功")
		return
	}

	if args.Install {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("安装服务成功")
		return
	}

	err = s.Run()
	if err != nil {
		_ = logger.Error(err)
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	//log.Println("===开始服务===")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	//log.Println("===停止服务===")
	_ = shutdown()
	return nil
}

func (p *Program) run() {

	// 此处编写具体的服务代码
	hup := make(chan os.Signal, 2)
	signal.Notify(hup, syscall.SIGHUP)
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		for {
			select {
			case <-hup:
			case <-quit:
				//优雅地结束
				_ = shutdown()
				//os.Exit(0)
			}
		}
	}()

	//原本的Main函数
	originMain()
}

func originMain() {
	banner.Print()
	build.Println()

	//加载主程序
	err := internal.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer internal.Close()

	//Web服务
	engine := web.CreateEngine()

	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup)

	//使用$前缀区分插件
	engine.Any("/app/:app/*path", app.ProxyApp)

	//监听Websocket
	engine.GET("/mqtt", broker.GinHandler)

	//前端静态文件
	engine.RegisterFS(http.FS(wwwFiles), "www", "index.html")

	//监听HTTP
	engine.Serve()
}

func shutdown() error {

	internal.Close()

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
