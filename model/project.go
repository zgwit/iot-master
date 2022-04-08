package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

type Template struct {
	ID      int    `json:"id" storm:"id,increment"`
	UUID    string `json:"uuid,omitempty"`
	Name    string `json:"name"`
	Version string `json:"version"` //SEMVER

	ProjectContent `storm:"inline"`

	Created time.Time `json:"created" storm:"created"`
}

type ProjectContent struct {
	Thumbnail   string        `json:"thumbnail"`
	Aggregators []*Aggregator `json:"aggregators"`
	Commands    []*Command    `json:"commands"`
	Jobs        []*Job        `json:"jobs"`
	Strategies  []*Strategy   `json:"strategies"`
}

//Project 项目
type Project struct {
	ID   int    `json:"id" storm:"id,increment"`
	Name string `json:"name"`

	Devices []*ProjectDevice `json:"devices"`
	//Devices []int `json:"devices"`

	TemplateId     int `json:"template_id,omitempty"`
	ProjectContent `storm:"inline"`

	Context calc.Context `json:"context"`

	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" storm:"created"`
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
	Created   time.Time `json:"created" storm:"created"`
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
