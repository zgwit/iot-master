package aggregator

import "github.com/robfig/cron/v3"

var _cron *cron.Cron

func Start() {
	_cron = cron.New()
	_cron.Start()
}
