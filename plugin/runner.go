package plugin

import (
	"github.com/zgwit/iot-master/v4/log"
	"os/exec"
	"time"
)

type Runner struct {
	path string
	cmd  *exec.Cmd
}

func (r *Runner) wait() {
	err := r.cmd.Wait()
	if err != nil {
		log.Error(err)
	}

	//重启 TODO 间隔可配置
	time.AfterFunc(time.Minute, func() {
		err := r.Start()
		if err != nil {
			log.Error(err)
		}
	})
}

func (r *Runner) Start() error {
	r.cmd = exec.Command(r.path)
	err := r.cmd.Start()
	if err != nil {
		return err
	}

	//协程里等待结果
	go r.wait()

	return nil
}
