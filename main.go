package main

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/kardianos/service"
	"github.com/zgwit/iot-master/v3/args"
	"github.com/zgwit/iot-master/v3/broker"
	"github.com/zgwit/iot-master/v3/config"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/vconn"
	"github.com/zgwit/iot-master/v3/web"
	"net"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

const banner = `
###        #######       #     #                                   
 #            #          ##   ##   ##    ####  ##### ###### #####  
 #    ####    #          # # # #  #  #  #        #   #      #    # 
 #   #    #   #   #####  #  #  # #    #  ####    #   #####  #    # 
 #   #    #   #          #     # ######      #   #   #      #####  
 #    ####    #          #     # #    # #    #   #   #      #   #  
###           #          #     # #    #  ####    #   ###### #    # 

`

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
	fmt.Print(banner)

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = log.Open(config.Config.Log)
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open(config.Config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(model.User), new(model.Password),
		new(model.Broker), new(model.Product),
		new(model.Device), new(model.Group), new(model.Alarm),
		new(model.App), new(model.Plugin),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = broker.Open(config.Config.Broker)
	if err != nil {
		return
	}
	defer broker.Close()

	if broker.Server != nil {
		err := mqtt.OpenBy(
			func(uri *url.URL, options paho.ClientOptions) (net.Conn, error) {
				c1, c2 := vconn.New()
				//EstablishConnection会读取connect，导致拥堵
				go func() {
					err := broker.Server.EstablishConnection("internal", c1)
					if err != nil {
						log.Error(err)
					}
				}()
				return c2, nil
			})
		if err != nil {
			return
		}

	} else {
		//MQTT总线
		err = mqtt.Open(config.Config.Mqtt)
		if err != nil {
			log.Fatal(err)
		}
		defer mqtt.Close()
	}

	//加载主程序
	err = internal.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer internal.Close()

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
