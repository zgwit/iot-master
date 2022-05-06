package model

import "time"

type Event struct {
	Id       int       `json:"id" storm:"id,increment"`
	UserId   int       `json:"user_id"`
	Event    string    `json:"event"`
	Target   string    `json:"target"`
	TargetId int       `json:"target_id"`
	Created  time.Time `json:"created" storm:"created"`
}
