package main

import (
	"github.com/kardianos/service"
	master "github.com/zgwit/iot-master/v4"
	"github.com/zgwit/iot-master/v4/args"
	"github.com/zgwit/iot-master/v4/banner"
	"github.com/zgwit/iot-master/v4/build"
	_ "github.com/zgwit/iot-master/v4/docs"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/web"
	"os"
	"os/signal"
	"syscall"
)

var serviceConfig = &service.Config{
	Name:        lib.AppName(),
	DisplayName: "物联大师",
	Description: "物联网数据中台",
	Arguments:   nil,
}

// @title 物联大师接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @InfoInstanceName master
// @query.collection.format multi
func main() {
	banner.Print()
	build.Println()

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
	_ = master.Shutdown()
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
				_ = master.Shutdown()
				//os.Exit(0)
			}
		}
	}()

	//原本的Main函数
	engine := web.CreateEngine()

	//启动
	err := master.Startup(engine)
	if err != nil {
		log.Fatal(err)
		//return
	}

	//注册静态页面
	fs := engine.FileSystem()
	master.Static(fs)

	//启动
	engine.Serve()
}
