package model

import (
	"time"
)

//Project 项目
type Project struct {
	Id   int64 `json:"id"`
	Name string `json:"name"`

	Devices []*ProjectDevice `json:"devices"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"` //编程名
}
