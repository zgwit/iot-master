package internal

import (
	"time"
)

type Device struct {
	Id         string
	Online     bool
	Last       time.Time
	Properties map[string]any
}

func NewDevice(id string) *Device {
	return &Device{
		Id:         id,
		Properties: make(map[string]any),
	}
}
