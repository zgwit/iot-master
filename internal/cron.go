package interval

import (
	"github.com/go-co-op/gocron"
	"time"
)

var Scheduler *gocron.Scheduler

func init()  {
	Scheduler = gocron.NewScheduler(time.UTC)
}
