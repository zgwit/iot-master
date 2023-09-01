package plugin

import (
	"fmt"
	db2 "github.com/zgwit/iot-master/v4/db"
	log2 "github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/model"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/pkg/lib"
	"github.com/zgwit/iot-master/v4/web"
	"os"
	"runtime"
)

var port = 40000

func getPort() int {
	port++
	if port >= 65535 {
		port = 40000
	}
	return port
}

type Plugin struct {
	*model.Plugin

	stop    bool
	Process *os.Process
}

func (p *Plugin) generateEnv(addr string) []string {
	ret := os.Environ()

	l := log2.GetOptions()
	s := l.ToEnv()
	ret = append(ret, s...)

	d := db2.GetOptions()
	s = d.ToEnv()
	ret = append(ret, s...)

	m := mqtt.GetOptions()
	m.ClientId = p.Id
	m.Username = p.Username
	m.Password = p.Password
	s = m.ToEnv()
	ret = append(ret, s...)

	w := web.GetOptions()
	w.Addr = addr

	s = w.ToEnv()
	ret = append(ret, s...)

	return ret
}

func (p *Plugin) Start() error {
	var err error

	//TODO linux下使用unix-sock
	addr := fmt.Sprintf(":%d", getPort())
	env := p.generateEnv(addr)

	cmd := p.Command

	//TODO 指定plugins目录，例如：plugins/alarm/alarm.exe
	if runtime.GOOS == "windows" {
		cmd = cmd + ".exe"
	} else {
		cmd = "./" + cmd
	}

	p.Process, err = os.StartProcess(cmd, []string{p.Id}, &os.ProcAttr{
		Files: []*os.File{nil, os.Stdout, os.Stderr}, //可以输出到日志文件
		Env:   env,
	})
	if err != nil {
		return err
	}
	p.Running = true

	//等待结束
	go func() {
		state, err := p.Process.Wait()
		p.Running = false
		log2.Info(state.ExitCode(), err)

		//异常退出，重新启动
		if p.stop {
			return
		}

		err = p.Start()
		if err != nil {
			log2.Error(err)
			return
		}
	}()

	return nil
}

func (p *Plugin) Close() error {
	p.stop = true
	return p.Process.Kill()
}

var plugins lib.Map[Plugin]

func New(model *model.Plugin) *Plugin {
	return &Plugin{
		Plugin: model,
		//Values: map[string]float64{},
	}
}

func Ensure(id string) (*Plugin, error) {
	p := plugins.Load(id)
	if p == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		p = plugins.Load(id)
	}
	return p, nil
}

func Get(id string) *Plugin {
	return plugins.Load(id)
}

func Load(id string) error {
	var p model.Plugin
	get, err := db2.Engine.ID(id).Get(&p)
	if err != nil {
		return err
	}
	if !get {
		return fmt.Errorf("plugin %s not found", id)
	}

	return From(&p)
}

func From(model *model.Plugin) error {
	p := New(model)

	err := p.Start()
	if err != nil {
		return err
	}

	plugins.Store(model.Id, p)

	return nil
}

func LoadAll() error {
	//开机加载所有插件
	var ps []*model.Plugin
	err := db2.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		if p.Disabled || p.External {
			continue
		}
		err = From(p)
		if err != nil {
			log2.Error(err)
			//return err
		}
	}

	return nil
}

func Close() {
	plugins.Range(func(id string, plugin *Plugin) bool {
		_ = plugin.Close()
		return true
	})
}
