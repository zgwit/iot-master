package model

import "time"

type Visit struct {
	Id       int       `json:"id" storm:"id,increment"`
	UserId   int       `json:"user_id" storm:"index"`
	Target   string    `json:"target"` //device project hmi
	TargetId int       `json:"target_id" storm:"index"`
	Last     time.Time `json:"last" storm:"index"`
	//Created  time.Time `json:"created" storm:"created"`
}
//TODO 更好的办法是栈式，比如限位20个，依次填入
