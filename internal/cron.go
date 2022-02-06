package interval

import (
	"github.com/go-co-op/gocron"
	"time"
)

var Cron *gocron.Scheduler

func init()  {
	Cron = gocron.NewScheduler(time.UTC)
}
