package model

import (
	"time"
)

type Device struct {
	Id         string             `json:"id" xorm:"pk"`                      //ClientID
	GatewayId  string             `json:"gateway_id,omitempty" xorm:"index"` //父设备
	ProductId  string             `json:"product_id,omitempty" xorm:"index"`
	GroupId    int64              `json:"group_id,omitempty" xorm:"index"` //分组
	Type       string             `json:"type,omitempty"`                  //网关/设备/子设备 gateway device subset
	Name       string             `json:"name"`
	Desc       string             `json:"desc,omitempty"`
	Username   string             `json:"username,omitempty"`
	Password   string             `json:"password,omitempty"`
	Parameters map[string]float64 `json:"parameters,omitempty"` //模型参数，用于报警检查
	Disabled   bool               `json:"disabled,omitempty"`
	Created    time.Time          `json:"created,omitempty" xorm:"created"`
}

type DeviceHistory struct {
	Id       int64     `json:"id" xorm:"pk"`
	DeviceId string    `json:"device_id" xorm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" xorm:"created"`
}
