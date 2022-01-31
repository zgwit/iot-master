package interval

import "time"

type Job struct {
	Disabled bool
	Type string //clock, crontab

	Clock int
	Crontab string

	WeekRanges []time.Weekday

	Invokes []Invoke
}


func (j *Job) Start() error {
	return nil
}


func (j *Job) Stop() error {
	return nil
}