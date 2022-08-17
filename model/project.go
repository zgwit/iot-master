package model

import (
	"time"
)

//Project 项目
type Project struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Devices []*ProjectDevice `json:"devices"`
	Jobs    []*Job           `json:"jobs"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   int64  `json:"id"`
	Name string `json:"name"` //编程名
}
