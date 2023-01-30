package model

import (
	"time"
)

// Device 设备
type Device struct {
	Id        string    `json:"id" xorm:"pk"`
	GatewayId string    `json:"gateway_id" xorm:"index"`
	ModelId   string    `json:"model_id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Disabled  bool      `json:"disabled"`
	Created   time.Time `json:"created" xorm:"created"`
}
