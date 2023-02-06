package model

import (
	"time"
)

type Server struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Port int    `json:"port"`
	//TODO 添加TLS证书
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}

type ServerHistory struct {
	Id       int64     `json:"id"`
	ServerId int64     `json:"server_id"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" xorm:"created"`
}
