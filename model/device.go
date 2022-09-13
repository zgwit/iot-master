package model

import (
	"time"
)

// Device 设备
type Device struct {
	Id        string `json:"id" xorm:"pk"`
	TunnelId  string `json:"tunnel_id" boltholdIndex:"TunnelId"`
	ProductId string `json:"product_id"`

	Name    string `json:"name"`
	Station int    `json:"station"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}
