package interval

import (
	"github.com/go-co-op/gocron"
	"time"
)

var Scheduler *gocron.Scheduler

func init()  {
	Scheduler = gocron.NewScheduler(time.UTC)
}
//TODO 是否是使用单一协程？？？ 是则要改成协程池？？？
//采集数据要等待，执行指令要下发数据，
