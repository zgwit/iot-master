package interval

import (
	"github.com/asaskevich/EventBus"
	"github.com/go-co-op/gocron"
	"time"
)

type Job struct {
	Disabled bool   `json:"disabled"`
	Type     string `json:"type"` //clock, crontab

	Clock   int    `json:"clock,omitempty"`
	Crontab string `json:"crontab,omitempty"`

	Weekdays []time.Weekday `json:"weekdays"`

	Invokes []Invoke `json:"invokes"`

	job *gocron.Job
	events EventBus.Bus
}

func (j *Job) Init() {
	j.events = EventBus.New()
}

func (j *Job) Start() error {
	var err error

	switch j.Type {
	case "clock":
		hours := j.Clock / 60
		minutes := j.Clock % 60
		j.job, err = Scheduler.At(hours).Hours().At(minutes).Minutes().Do(func() {
			j.Execute()
		})
	case "crontab":
		j.job, err = Scheduler.Cron(j.Crontab).Do(func() {
			j.Execute()
		})
	}
	return err
}

func (j *Job) Execute() {
	//for _, i:= range j.Invokes {
	//	j.events.Publish("invoke", i)
	//}
	//避免拥堵计时器
	go j.events.Publish("invoke")
}

func (j *Job) Stop() {
	Scheduler.Remove(j.job)
}
