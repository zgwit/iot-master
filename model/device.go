package model

import (
	"time"
)

type DeviceHistory struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	History  string    `json:"history"`
	Created  time.Time `json:"created"`
}

type DeviceHistoryAlarm struct {
	Id int `json:"id" storm:"id,increment"`

	DeviceId int    `json:"Device_id"`
	Code     string `json:"code"`
	Level    int    `json:"level"`
	Message  string `json:"message"`

	Created time.Time `json:"created"`
}

type DeviceHistoryJob struct {
	Id      int       `json:"id" storm:"id,increment"`
	Job     string    `json:"job"`
	Result  string    `json:"result"`
	Created time.Time `json:"created"`
}

type DeviceHistoryCommand struct {
	Id      int       `json:"id" storm:"id,increment"`
	Command string    `json:"command"`
	Argv    string    `json:"argv"`
	Result  string    `json:"result"`
	Created time.Time `json:"created"`
}
