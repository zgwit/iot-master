package model

import (
	"time"
)

//Product 产品
type Product struct {
	Id           string   `json:"id" xorm:"pk"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"` //厂家
	Version      string   `json:"version"`      //SEMVER
	Protocol     Protocol `json:"protocol" xorm:"JSON"`
	//Tunnel       string `json:"tunnel"` // serial tcp udp ???

	Tags     []string   `json:"tags,omitempty"`
	Points   []*Point   `json:"points"`
	Pollers  []*Poller  `json:"pollers"`
	Commands []*Command `json:"commands"`

	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
}

//Device 设备
type Device struct {
	Id        int64  `json:"id"`
	TunnelId  int64  `json:"tunnel_id" xorm:"index"`
	ProductId string `json:"product_id"`

	Name    string `json:"name"`
	Station int    `json:"station"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
}
