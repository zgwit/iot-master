package model

import (
	"time"
)

type ProjectHistory struct {
	Id        int
	ProjectId int
	History   string
	Created   time.Time
}

type ProjectHistoryAlarm struct {
	Id int

	ProjectId int    `json:"project_id"`
	DeviceId  int    `json:"device_id"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
	Message   string `json:"message"`

	Created time.Time
}

type ProjectHistoryJob struct {
	Id      int
	Job     string
	Result  string
	Created time.Time
}
