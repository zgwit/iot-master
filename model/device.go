package model

import (
	"time"
)

type Device struct {
	Id         string             `json:"id" xorm:"pk"` //ClientID
	ProductId  string             `json:"product_id" xorm:"index"`
	ParentId   string             `json:"parent_id" xorm:"index"` //父设备
	Type       string             `json:"type"`                   //网关/设备/子设备 gateway device subset
	Name       string             `json:"name"`
	Desc       string             `json:"desc"`
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Parameters map[string]float64 `json:"parameters"` //模型参数，用于报警检查
	Disabled   bool               `json:"disabled"`
	Created    time.Time          `json:"created" xorm:"created"`
}

type DeviceHistory struct {
	Id       int64     `json:"id" xorm:"pk"`
	DeviceId string    `json:"device_id" xorm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" xorm:"created"`
}
