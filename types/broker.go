package types

import "time"

type Broker struct {
	Id       string    `json:"id" xorm:"pk"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc,omitempty"`
	Port     int       `json:"port,omitempty"` //TODO 添加TLS证书
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}

type BrokerHistory struct {
	Id       int64     `json:"id"`
	BrokerId string    `json:"broker_id" xorm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}
