package model

import (
	"time"
)

//UserJob 用户定时任务
type UserJob struct {
	ID       int  `json:"id" storm:"id,increment"`
	UserID   int  `json:"user_id"`
	Disabled bool `json:"disabled"`

	Clock    int            `json:"clock"`
	Weekdays []time.Weekday `json:"weekdays"`

	Invokes []Invoke `json:"invokes"`
}
