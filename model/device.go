package model

import (
	"time"
)

//Element 元件
type Element struct {
	Id           string `json:"id" xorm:"pk"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"` //厂家
	Version      string `json:"version"`      //SEMVER

	DeviceContent `xorm:"extends"`

	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
	Deleted time.Time `json:"-" xorm:"deleted"`
}

type DeviceContent struct {
	HMI         string        `json:"hmi" xorm:"'hmi'"`
	Tags        []string      `json:"tags,omitempty"`
	Points      []*Point      `json:"points"`
	Pollers     []*Poller     `json:"pollers"`
	Calculators []*Calculator `json:"calculators"`
	Commands    []*Command    `json:"commands"`
	Alarms      []*Alarm      `json:"alarms"`
}

//Device 设备
type Device struct {
	Id        int64  `json:"id"`
	TunnelId  int64  `json:"tunnel_id" xorm:"index"`
	ElementId string `json:"element_id"`

	Name          string `json:"name"`
	Station       int    `json:"station"`
	DeviceContent `xorm:"extends"`

	//上下文
	//Context map[string]interface{} `json:"context"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	Deleted  time.Time `json:"-" xorm:"deleted"`
}

type DeviceEx struct {
	Device  `xorm:"extends"`
	Running bool   `json:"running"`
	Tunnel  string `json:"tunnel"`
	Element string `json:"element"`
}

func (d *DeviceEx) TableName() string {
	return "device"
}

type DeviceHistory struct {
	Device   `xorm:"extends"`
	DeviceId int64 `json:"device_id" xorm:"index"`
}

type ElementHistory struct {
	Element   `xorm:"extends"`
	ElementId string `json:"element_id" xorm:"index"`
}
