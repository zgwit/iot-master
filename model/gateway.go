package model

import "time"

type Gateway struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Created time.Time `json:"created" xorm:"created"`
}
