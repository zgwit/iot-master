package service

import (
	"github.com/kardianos/service"
	"github.com/zgwit/iot-master/v4/pkg/config"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

var svc service.Service

func Register(startup, shutdown func() error) (err error) {
	var serviceConfig = &service.Config{
		Name:        config.GetString(MODULE, "name"),
		DisplayName: config.GetString(MODULE, "display"),
		Description: config.GetString(MODULE, "description"),
		Arguments:   config.GetStringSlice(MODULE, "arguments"),
	}

	p := &Program{
		Startup:  startup,
		Shutdown: shutdown,
	}

	svc, err = service.New(p, serviceConfig)
	return
}

func Run() error {
	return svc.Run()
}

func Start() error {
	return svc.Start()
}

func Restart() error {
	return svc.Restart()
}

func Install() error {
	return svc.Install()
}

func Uninstall() error {
	return svc.Install()
}

type Program struct {
	Startup  func() error
	Shutdown func() error
}

func (p *Program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	return p.Shutdown()
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
				_ = p.Shutdown()
				//os.Exit(0)
			}
		}
	}()

	//内部启动
	err := p.Startup()
	if err != nil {
		log.Error(err)
	}
}
