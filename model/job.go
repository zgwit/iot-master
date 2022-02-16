package model

import (
	"time"
)

//Job 任务
type Job struct {
	Disabled bool `json:"disabled"`

	Type string `json:"type"` //clock, crontab

	Clock    int            `json:"clock,omitempty"`
	Weekdays []time.Weekday `json:"weekdays"`

	Crontab string `json:"crontab,omitempty"`

	Invokes []Invoke `json:"invokes"`
}
