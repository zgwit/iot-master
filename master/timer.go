package master

import (
	"fmt"
	"github.com/zgwit/iot-master/cron"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
)

//Timer 任务
type Timer struct {
	model.Timer

	job *cron.Job

	events.EventEmitter
}

//Start 启动任务
func (j *Timer) Start() error {
	var err error
	hours := j.Clock / 60
	minutes := j.Clock % 60
	//TODO 处理weekdays
	j.job, err = cron.Clock(hours, minutes, func() {
		j.Execute()
	})
	return err
}

//Execute 执行任务
func (j *Timer) Execute() {
	//for _, i:= range j.Invokes {
	//	j.events.Publish("invoke", i)
	//}
	//TODO 避免拥堵计时器(需要确认)
	go j.Emit("invoke")
}

//Stop 取消任务
func (j *Timer) Stop() {
	j.job.Cancel()
}

//String 任务描述
func (j *Timer) String() string {
	hours := j.Clock / 60
	minutes := j.Clock % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
