package plugin

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"os"
)

type Plugin struct {
	*model.Plugin

	Process *os.Process
}

func (p *Plugin) Start() error {
	var err error

	env := os.Environ()
	//TODO 加入初始化配置，日志等级，端口号等

	p.Process, err = os.StartProcess(p.Command, nil, &os.ProcAttr{
		Files: nil, //TODO 输出到日志文件
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
		if err != nil {
			log.Error(err)
			return
		}
		log.Info(state.ExitCode())
	}()

	return nil
}

func (p *Plugin) Close() error {
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
	get, err := db.Engine.ID(id).Get(&p)
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

	plugins.Store(model.Id, p)

	return nil
}

func LoadAll() error {
	//开机加载所有插件，好像没有必要???
	var ps []*model.Plugin
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = From(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
