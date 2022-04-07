package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

//Element 元件
type Element struct {
	ID      int    `json:"id" storm:"id,increment"`
	UUID    string `json:"uuid,omitempty"`
	Name    string `json:"name"`
	Version string `json:"version"` //SEMVER

	DeviceContent `storm:"inline"`

	Created time.Time `json:"created" storm:"created"`
}

type DeviceContent struct {
	Tags []string `json:"tags,omitempty"`

	//从机号
	Mapper *Mapping `json:"mapper"` //内存映射

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
	DeviceContent `storm:"inline"`

	//上下文
	Context calc.Context `json:"context"`

	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" storm:"created"`
}

//DeviceHistory 设备历史
type DeviceHistory struct {
	ID       int       `json:"id" storm:"id,increment"`
	DeviceID int       `json:"device_id" storm:"index"`
	History  string    `json:"history"`
	Created  time.Time `json:"created" storm:"created"`
}

//DeviceHistoryAlarm 设备历史告警
type DeviceHistoryAlarm struct {
	DeviceHistory `storm:"inline"`
	Code          string `json:"code"`
	Level         int    `json:"level"`
	Message       string `json:"message"`
}

//DeviceHistoryReactor 设备历史响应
type DeviceHistoryReactor struct {
	DeviceHistory `storm:"inline"`
	Name          string `json:"name"`
}

//DeviceHistoryJob 设备历史任务
type DeviceHistoryJob struct {
	DeviceHistory `storm:"inline"`
	Job           string `json:"job"`
}

//DeviceHistoryCommand 设备历史命令
type DeviceHistoryCommand struct {
	DeviceHistory `storm:"inline"`
	Command       string    `json:"command"`
	Argv          []float64 `json:"argv"`
}

//DeviceHistoryTimer 设备定时任务
type DeviceHistoryTimer struct {
	DeviceHistory `storm:"inline"`
	TimerID       int `json:"timer_id" storm:"index"`
}
