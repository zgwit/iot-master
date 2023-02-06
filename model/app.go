package model

import "time"

type App struct {
	Name    string `json:"name"`
	Icon    string `json:"icon,omitempty"`
	Label   string `json:"label"`
	Desc    string `json:"desc"`
	Type    string `json:"type"` //tcp unix
	Address string `json:"address"`
	Hidden  bool   `json:"hidden,omitempty"` //隐藏，适用于服务
}

type AppHistory struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Event   string    `json:"event"`
	Created time.Time `json:"created" xorm:"created"`
}
