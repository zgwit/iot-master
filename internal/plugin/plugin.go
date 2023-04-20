package plugin

import (
	"github.com/zgwit/iot-master/v3/model"
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
