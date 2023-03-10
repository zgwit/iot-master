package model

import "time"

type PayloadPropertyUp struct {
	PayloadDevice
	//子设备
	Devices []PayloadDevice `json:"devices,omitempty"`
}

type PayloadValue struct {
	Name      string    `json:"name"`
	Time      time.Time `json:"time,omitempty"`
	Timestamp int64     `json:"timestamp,omitempty"`
	Value     any       `json:"value"`
}

type PayloadDevice struct {
	Id         string         `json:"id"`
	Time       time.Time      `json:"time,omitempty"`
	Timestamp  int64          `json:"timestamp,omitempty"`
	Properties []PayloadValue `json:"properties"`
}

type PayloadEvent struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Title   string         `json:"title"`
	Message string         `json:"message,omitempty"`
	Output  map[string]any `json:"output"`
}
