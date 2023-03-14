package model

import (
	"time"
)

type Broker struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Port     int       `json:"port"` //TODO 添加TLS证书
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}

type BrokerHistory struct {
	Id       int64     `json:"id"`
	BrokerId int64     `json:"broker_id" xorm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" xorm:"created"`
}
