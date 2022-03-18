package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

//Project 项目
type Project struct {
	ID       int    `json:"id" storm:"id,increment"`
	Disabled bool   `json:"disabled,omitempty"`
	Template string `json:"template,omitempty"`

	Devices []*ProjectDevice `json:"devices"`
	//Devices []int `json:"devices"`

	Aggregators []*Aggregator `json:"aggregators"`
	Commands    []*Command    `json:"commands"`
	Jobs        []*Job        `json:"jobs"`
	Strategies  []*Strategy   `json:"strategies"`

	Context calc.Context `json:"context"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	ID   int    `json:"id"`
	Name string `json:"name"` //编程名
}

//ProjectHistory 项目历史
type ProjectHistory struct {
	ID        int       `json:"id" storm:"id,increment"`
	ProjectID int       `json:"project_id"`
	History   string    `json:"history"`
	Created   time.Time `json:"created"`
}

//ProjectHistoryAlarm 项目历史告警
type ProjectHistoryAlarm struct {
	ProjectHistory `storm:"inline"`
	ProjectAlarm   `storm:"inline"`
}

//ProjectHistoryStrategy 项目历史响应
type ProjectHistoryStrategy struct {
	ProjectHistory `storm:"inline"`
	Name           string `json:"name"`
}

//ProjectHistoryJob 项目历史任务
type ProjectHistoryJob struct {
	ProjectHistory `storm:"inline"`
	Job            string `json:"job"`
}

//ProjectHistoryTimer 设备定时任务
type ProjectHistoryTimer struct {
	ProjectHistory `storm:"inline"`
	TimerID        int `json:"timer_id" storm:"index"`
}
