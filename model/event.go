package model

import "time"

type Event struct {
	Id       int64     `json:"id" xorm:"autoincr"`
	UserId   int64     `json:"user_id"`
	Event    string    `json:"event"`
	Target   string    `json:"target"`
	TargetId int64     `json:"target_id"`
	Created  time.Time `json:"created" xorm:"created"`
}
