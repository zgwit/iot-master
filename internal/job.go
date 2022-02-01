package interval

import "time"

type Job struct {
	Disabled bool   `json:"disabled"`
	Type     string `json:"type"` //clock, crontab

	Clock   int    `json:"clock"`
	Crontab string `json:"crontab"`

	WeekRanges []time.Weekday `json:"week_ranges"`

	Invokes []Invoke `json:"invokes"`
}

func (j *Job) Start() error {
	return nil
}

func (j *Job) Stop() error {
	return nil
}
