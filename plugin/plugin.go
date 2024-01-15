package plugin

import (
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/log"
	"os"
	"path/filepath"
	"runtime"
	"time"
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

	if p.Main == "" {
		return nil
	}

	cmd := filepath.Join(viper.GetString("data"), "plugin", p.Id, p.Main)
	dir := filepath.Join(viper.GetString("data"), "plugin", p.Id)

	//绝对路径
	cmd, err = filepath.Abs(cmd)
	if err != nil {
		return err
	}

	//工作目录
	//dir := filepath.Dir(cmd)

	//TODO 指定plugins目录，例如：plugins/alarm/alarm.exe
	if runtime.GOOS == "windows" {
		cmd = cmd + ".exe"
	}

	p.Process, err = os.StartProcess(cmd, []string{p.Id}, &os.ProcAttr{
		Dir:   dir,                                   //使用插件目录
		Files: []*os.File{nil, os.Stdout, os.Stderr}, //TODO 输出到日志文件
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

		//应该做成智能休眠
		time.Sleep(time.Second * 5)

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
