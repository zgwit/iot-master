package internal

import (
	"github.com/go-co-op/gocron"
)

type Collector struct {
	Disabled bool   `json:"disabled"`
	Type     string `json:"type"` //interval, clock, crontab
	Interval int    `json:"interval,omitempty"`
	Clock    int    `json:"clock,omitempty"`
	Crontab  string `json:"crontab,omitempty"`

	Code    int `json:"code"`
	Address int `json:"address"`
	//TODO Address2
	Length int `json:"length"`

	//TODO Filters

	//等待结果
	Parallel bool `json:"parallel,omitempty"`

	reading bool
	job     *gocron.Job
	adapter *Adapter
}

func (c *Collector) Start() (err error) {
	switch c.Type {
	case "interval":
		c.job, err = Scheduler.Every(c.Interval).Milliseconds().Do(func() {
			c.Execute()
		})
	case "clock":
		hours := c.Clock / 60
		minutes := c.Clock % 60
		c.job, err = Scheduler.At(hours).Hours().At(minutes).Minutes().Do(func() {
			c.Execute()
		})
	case "crontab":
		c.job, err = Scheduler.Cron(c.Crontab).Do(func() {
			c.Execute()
		})
	}
	return
}

func (c *Collector) Execute() {
	//阻塞情况下，采集数据，要等待，避免大量读指令阻塞
	if !c.Parallel && c.reading {
		return
	}

	//TODO 此举会不断创建协程 需要再确定gocron的协程机制
	go c.read()
}

func (c *Collector) read()  {
	c.reading = true
	_, _ = c.adapter.Read(c.Code, c.Address, c.Length)
	//log error
	c.reading = false
}

func (c *Collector) Stop() {
	Scheduler.Remove(c.job)
}
