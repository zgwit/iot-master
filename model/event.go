package model

import "time"

type Event struct {
	Id       int64          `json:"id"`
	DeviceId string         `json:"device_id" xorm:"index"`
	Name     string         `json:"name"`
	Output   map[string]any `json:"output"`
	Created  time.Time      `json:"created" xorm:"created"`
}

type ModEvent struct {
	Name   string        `json:"name"`
	Label  string        `json:"label"`
	Desc   string        `json:"desc,omitempty"`
	Type   string        `json:"type"` //info alert error //Level  uint8         `json:"level"`
	Output []ModArgument `json:"output"`
}
