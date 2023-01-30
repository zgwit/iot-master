package model

import "time"

type Gateway struct {
	Id       string    `json:"id" xorm:"pk"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Client   string    `json:"client,omitempty"` //ClientID 是否与ID重复？
	Username string    `json:"username"`
	Password string    `json:"password"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}

type GatewayHistory struct {
	Id        int64     `json:"id" xorm:"pk"`
	GatewayId string    `json:"gateway_id" xorm:"index"`
	Event     string    `json:"event"`
	Created   time.Time `json:"created" xorm:"created"`
}
