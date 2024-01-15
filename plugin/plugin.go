package plugin

import (
	"github.com/zgwit/iot-master/v4/log"
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
	*Manifest

	Addr    string
	Running bool

	stop    bool
	Process *os.Process
}

func (p *Plugin) Start() error {
	var err error

	//TODO linux下使用unix-sock
	//addr := fmt.Sprintf(":%d", getPort())
	//env := p.generateEnv(addr)

	cmd := p.Main //TODO 插件目录

	//TODO 指定plugins目录，例如：plugins/alarm/alarm.exe
	if runtime.GOOS == "windows" {
		cmd = cmd + ".exe"
	} else {
		cmd = "./" + cmd
	}

	p.Process, err = os.StartProcess(cmd, []string{p.Id}, &os.ProcAttr{
		Dir:   ".",                                   //使用插件目录
		Files: []*os.File{nil, os.Stdout, os.Stderr}, //可以输出到日志文件
		//Env:   env,
	})
	if err != nil {
		return err
	}
	p.Running = true

	//等待结束
	go func() {
		state, err := p.Process.Wait()
		p.Running = false
		log.Info(state.ExitCode(), err)

		//异常退出，重新启动
		if p.stop {
			return
		}

		err = p.Start()
		if err != nil {
			log.Error(err)
			return
		}
	}()

	return nil
}

func (p *Plugin) Close() error {
	p.stop = true
	return p.Process.Kill()
}
