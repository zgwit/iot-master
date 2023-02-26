package model

import (
	"time"
)

type Device struct {
	Id        string    `json:"id" xorm:"pk"` //ClientID
	ProductId string    `json:"product_id"`
	Gateway   bool      `json:"gateway"` //网关类型，否为直连设备
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Disabled  bool      `json:"disabled"`
	Created   time.Time `json:"created" xorm:"created"`
}

type DeviceHistory struct {
	Id       int64     `json:"id" xorm:"pk"`
	DeviceId string    `json:"device_id" xorm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" xorm:"created"`
}

// Subset 设备
type Subset struct {
	Id        string    `json:"id" xorm:"pk"`
	DeviceId  string    `json:"device_id" xorm:"index"`
	ProductId string    `json:"product_id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Disabled  bool      `json:"disabled"`
	Created   time.Time `json:"created" xorm:"created"`
}
