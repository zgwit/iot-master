package common

import (
	"github.com/go-co-op/gocron"
	"time"
)

var Scheduler *gocron.Scheduler

func init()  {
	Scheduler = gocron.NewScheduler(time.UTC)
}
//是否是使用单一协程？？？ 是则要改成协程池？？？
