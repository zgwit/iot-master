package model

import (
	"time"
)

// Project 项目
type Project struct {
	Id          string `json:"id" xorm:"pk"`
	Name        string `json:"name"`
	InterfaceId string `json:"interface_id"`

	Devices []*ProjectDevice `json:"devices"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}

// ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"` //编程名
}
