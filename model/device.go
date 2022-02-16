package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

type Device struct {
	Id      int      `json:"id" storm:"id,increment"`
	Element string   `json:"element,omitempty"` //元件UUID
	Name    string   `json:"name"`
	Tags    []string `json:"tags,omitempty"`

	//从机号
	Slave int `json:"slave"`

	Points      []*Point      `json:"points"`
	Pollers     []*Poller     `json:"pollers"`
	Calculators []*Calculator `json:"calculators"`
	Commands    []*Command    `json:"commands"`
	Reactors    []*Rule       `json:"reactors"`
	Jobs        []*Job        `json:"jobs"`

	//上下文
	Context calc.Context `json:"context"`

	Disabled bool `json:"disabled,omitempty"`
}

//DeviceHistory 设备历史
type DeviceHistory struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	History  string    `json:"history"`
	Created  time.Time `json:"created"`
}

//DeviceHistoryAlarm 设备历史告警
type DeviceHistoryAlarm struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	Code     string    `json:"code"`
	Level    int       `json:"level"`
	Message  string    `json:"message"`
	Created  time.Time `json:"created"`
}

//DeviceHistoryReactor 设备历史响应
type DeviceHistoryReactor struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	Name     string    `json:"name"`
	History  string    `json:"result"`
	Created  time.Time `json:"created"`
}

//DeviceHistoryJob 设备历史任务
type DeviceHistoryJob struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	Job      string    `json:"job"`
	History  string    `json:"result"`
	Created  time.Time `json:"created"`
}

//DeviceHistoryCommand 设备历史命令
type DeviceHistoryCommand struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	Command  string    `json:"command"`
	Argv     []float64 `json:"argv"`
	History  string    `json:"result"`
	Created  time.Time `json:"created"`
}
