package master

import (
	"github.com/zgwit/iot-master/cron"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
)

//Poller 采集器
type Poller struct {
	model.Poller
	Addr protocol.Addr

	reading bool
	job     *cron.Job
	adapter *Mapper
}

//Start 启动
func (p *Poller) Start() (err error) {
	switch p.Type {
	case "interval":
		p.job, err = cron.Interval(p.Interval, func() {
			p.Execute()
		})
	case "clock":
		hours := p.Clock / 60
		minutes := p.Clock % 60
		p.job, err = cron.Clock(hours, minutes, func() {
			p.Execute()
		})
	case "crontab":
		p.job, err = cron.Schedule(p.Crontab, func() {
			p.Execute()
		})
	}
	return
}


//Execute 执行
func (p *Poller) Execute() {
	//阻塞情况下，采集数据，要等待，避免大量读指令阻塞
	//if !p.Parallel && p.reading {
	if p.reading {
		return
	}

	//TODO 此举会不断创建协程 需要再确定gocron的协程机制
	go p.read()
}

func (p *Poller) read() {
	p.reading = true
	_, _ = p.adapter.Read(p.Addr, p.Length)
	//log error
	p.reading = false
}

//Stop 结束
func (p *Poller) Stop() {
	p.job.Cancel()
}
