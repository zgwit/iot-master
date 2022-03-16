package model

import (
	"time"
)

//Timer 用户定时任务
type Timer struct {
	ID       int  `json:"id" storm:"id,increment"`
	UserID   int  `json:"user_id"`
	Disabled bool `json:"disabled"`

	Clock    int            `json:"clock"`
	Weekdays []time.Weekday `json:"weekdays"`

	Invokes []Invoke `json:"invokes"`
}

type ProjectTimer struct {
	Timer     `storm:"inline"`
	ProjectId int `json:"project_id" storm:"index"`
}

type DeviceTimer struct {
	Timer    `storm:"inline"`
	DeviceId int `json:"device_id" storm:"index"`
}
