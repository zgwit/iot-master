package master

import (
	"fmt"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/master/cron"
	"time"
)

type Job struct {
	Disabled bool   `json:"disabled"`
	Type     string `json:"type"` //clock, crontab

	Clock    int            `json:"clock,omitempty"`
	Weekdays []time.Weekday `json:"weekdays"`

	Crontab string `json:"crontab,omitempty"`

	Invokes []*Invoke `json:"invokes"`

	job *cron.Job

	events.EventEmitter
}

func (j *Job) Start() error {
	var err error

	switch j.Type {
	case "clock":
		hours := j.Clock / 60
		minutes := j.Clock % 60
		//TODO 处理weekdays
		j.job, err = cron.Clock(hours, minutes, func() {
			j.Execute()
		})
	case "crontab":
		j.job, err = cron.Schedule(j.Crontab, func() {
			j.Execute()
		})
	}
	return err
}

func (j *Job) Execute() {
	//for _, i:= range j.Invokes {
	//	j.events.Publish("invoke", i)
	//}
	//TODO 避免拥堵计时器(需要确认)
	go j.Emit("invoke")
}

func (j *Job) Stop() {
	j.job.Cancel()
}

func (j *Job) String() string {
	switch j.Type {
	case "clock":
		hours := j.Clock / 60
		minutes := j.Clock % 60
		return fmt.Sprintf("%02d:%02d", hours, minutes)
	case "crontab":
		return j.Crontab
	}
	return ""
}
