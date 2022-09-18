package model

import "time"

type Interface struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Created time.Time `json:"created" xorm:"created"`
}
