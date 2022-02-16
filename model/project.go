package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

//Project 项目
type Project struct {
	Id       int    `json:"id" storm:"id,increment"`
	Disabled bool   `json:"disabled,omitempty"`
	Template string `json:"template,omitempty"`

	Devices []*ProjectDevice `json:"devices"`
	//Devices []int `json:"devices"`

	Aggregators []*Aggregator `json:"aggregators"`
	Commands    []*Command    `json:"commands"`
	Reactors    []*Reactor    `json:"reactors"`
	Jobs        []*Job        `json:"jobs"`

	Context calc.Context `json:"context"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   int    `json:"id"`
	Name string `json:"name"` //编程名
}

//ProjectHistory 项目历史
type ProjectHistory struct {
	Id        int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	History   string    `json:"history"`
	Created   time.Time `json:"created"`
}

//ProjectHistoryAlarm 项目历史告警
type ProjectHistoryAlarm struct {
	Id int `json:"id" storm:"id,increment"`

	ProjectId int    `json:"project_id"`
	DeviceId  int    `json:"device_id"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
	Message   string `json:"message"`

	Created time.Time `json:"created"`
}

//ProjectHistoryReactor 项目历史响应
type ProjectHistoryReactor struct {
	Id        int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	Name      string    `json:"name"`
	History   string    `json:"result"`
	Created   time.Time `json:"created"`
}

//ProjectHistoryJob 项目历史任务
type ProjectHistoryJob struct {
	Id        int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	Job       string    `json:"job"`
	History   string    `json:"result"`
	Created   time.Time `json:"created"`
}
