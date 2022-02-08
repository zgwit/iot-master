package model

import (
	"time"
)

type ProjectHistory struct {
	Id        int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	History   string    `json:"history"`
	Created   time.Time `json:"created"`
}

type ProjectHistoryAlarm struct {
	Id int `json:"id" storm:"id,increment"`

	ProjectId int    `json:"project_id"`
	DeviceId  int    `json:"device_id"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
	Message   string `json:"message"`

	Created time.Time `json:"created"`
}

type ProjectHistoryJob struct {
	Id      int       `json:"id" storm:"id,increment"`
	Job     string    `json:"job"`
	Result  string    `json:"result"`
	Created time.Time `json:"created"`
}
