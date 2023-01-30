package model

import (
	"time"
)

type Server struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Port int    `json:"port"`
	//TODO 添加TLS证书
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}
