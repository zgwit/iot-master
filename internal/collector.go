package interval

import (
	"github.com/asaskevich/EventBus"
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

	job    *gocron.Job
	events EventBus.Bus
	adapter *Adapter
}

func (c *Collector) Init() {
	c.events = EventBus.New()
}

func (c *Collector) Start() (err error) {
	switch c.Type {
	case "interval":
		c.job, err = Scheduler.Every(1).Milliseconds().Do(func() {
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

func (c *Collector) Execute() error {
	//TODO 上报？？采集？？
	//c.events.Publish("action")
	data, err := c.adapter.Read(c.Code, c.Address, c.Length)
	if err != nil {
		return err
	}
	//上报
	c.events.Publish("data", data)
	return nil
}

func (c *Collector) Stop() {
	Scheduler.Remove(c.job)
}
