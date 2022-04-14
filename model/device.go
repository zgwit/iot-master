package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

//Element 元件
type Element struct {
	ID           int    `json:"id" storm:"id,increment"`
	UUID         string `json:"uuid,omitempty"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"` //厂家
	Version      string `json:"version"`      //SEMVER

	DeviceContent `storm:"inline"`

	Created time.Time `json:"created" storm:"created"`
}

type DeviceContent struct {
	Icon string   `json:"icon"`
	Tags []string `json:"tags,omitempty"`

	//从机号
	//Mapper *Mapping `json:"mapper"` //内存映射
	Points      []*Point      `json:"points"`
	Pollers     []*Poller     `json:"pollers"`
	Calculators []*Calculator `json:"calculators"`
	Commands    []*Command    `json:"commands"`
	Jobs        []*Job        `json:"jobs"`
	Strategies  []*Strategy   `json:"strategies"`
}

//Device 设备
type Device struct {
	ID        int `json:"id" storm:"id,increment"`
	LinkId    int `json:"link_id" storm:"index"`
	ElementId int `json:"element_id"`

	Name          string `json:"name"`
	Station       int    `json:"station"`
	DeviceContent `storm:"inline"`

	//上下文
	Context calc.Context `json:"context"`

	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" storm:"created"`
}

//DeviceEvent 设备事件
type DeviceEvent struct {
	ID       int       `json:"id" storm:"id,increment"`
	DeviceID int       `json:"device_id" storm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" storm:"created"`
}

type DeviceHistory struct {
	Device   `storm:"inline"`
	DeviceID int `json:"device_id" storm:"unique"`
}

type ElementHistory struct {
	Element   `storm:"inline"`
	ElementID int `json:"device_id" storm:"unique"`
}
