package model

import (
	"time"
)

type Device struct {
	Id        string    `json:"id" xorm:"pk"` //ClientID
	ProductId string    `json:"product_id"`
	DeviceId  string    `json:"device_id" xorm:"index"` //父设备
	IsGateway string    `json:"is_gateway"`             //网关设备，有mqtt用户名，密码
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
