package model

import (
	"time"
)

type DeviceHistory struct {
	Id       int
	DeviceId int
	History  string
	Created  time.Time
}

type DeviceHistoryAlarm struct {
	Id int

	DeviceId int    `json:"Device_id"`
	Code     string `json:"code"`
	Level    int    `json:"level"`
	Message  string `json:"message"`

	Created time.Time
}

type DeviceHistoryJob struct {
	Id      int
	Job     string
	Result  string
	Created time.Time
}

type DeviceHistoryCommand struct {
	Id      int
	Command string
	Argv    string
	Result  string
	Created time.Time
}
