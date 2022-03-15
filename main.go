package main

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/zgwit/iot-master/args"
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/mqtt"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var serviceConfig = &service.Config{
	Name:        "iot-master",
	DisplayName: "物联大师",
	Description: "物联网设备自动控制系统",
	Arguments: nil,
}

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

	if args.Install {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
		return
	}

	if args.Uninstall {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
		return
	}

	err = s.Run()
	if err != nil {
		_ = logger.Error(err)
	}
}


type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("===开始服务===")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("===停止服务===")
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
				//TODO 优雅地结束
				os.Exit(0)
			}
		}
	}()

	//原本的Main函数
	originMain()
}

func originMain() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Open(config.Config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	err = tsdb.Open(config.Config.History)
	if err != nil {
		log.Fatal(err)
	}
	defer tsdb.Close()

	err = connect.LoadTunnels()
	if err != nil {
		log.Fatal(err)
	}
	//defer connect.Close()

	err = master.Start()
	if err != nil {
		log.Fatal(err)
	}
	//defer master.Close()

	//TODO 正确创建 MQTT Broker
	mqtt.NewBroker()

	//判断是否开启Web
	web.Serve(config.Config.Web)
}

func shutdown() error {

	//_ = database.Close()
	//_ = tsdb.Close()
	//connect.Close()
	//master.Close()

	//只关闭Web就行了，其他通过defer关闭
	_ = web.Close()

	return nil
}