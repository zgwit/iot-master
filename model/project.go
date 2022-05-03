package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

type Template struct {
	Id      string `json:"id" storm:"id"`
	Name    string `json:"name"`
	Version string `json:"version"` //SEMVER

	Elements []*TemplateElement `json:"elements"`

	ProjectContent `storm:"inline"`

	Created time.Time `json:"created" storm:"created"`
}

type ProjectContent struct {
	Icon        string        `json:"icon"`
	Aggregators []*Aggregator `json:"aggregators"`
	Jobs        []*Job        `json:"jobs"`
	Alarms      []*Alarm      `json:"alarms"`
	Strategies  []*Strategy   `json:"strategies"`
	//Commands    []*Command    `json:"commands"`
}

//Project 项目
type Project struct {
	Id   int    `json:"id" storm:"id,increment"`
	Name string `json:"name"`

	Devices []*ProjectDevice `json:"devices"`
	//Devices []int `json:"devices"`

	TemplateId     string `json:"template_id,omitempty"`
	ProjectContent `storm:"inline"`

	Context calc.Context `json:"context"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" storm:"created"`
}

type ProjectEx struct {
	Project
	Running  bool   `json:"running"`
	Template string `json:"template"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   int    `json:"id"`
	Name string `json:"name"` //编程名
}

type TemplateElement struct {
	Id   string `json:"id"`
	Name string `json:"name"` //编程名
}

//ProjectEvent 项目历史
type ProjectEvent struct {
	Id        int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	Event     string    `json:"event"`
	Created   time.Time `json:"created" storm:"created"`
}

type ProjectHistory struct {
	Project   `storm:"inline"`
	ProjectId int `json:"project_id" storm:"index"`
}

type TemplateHistory struct {
	Template   `storm:"inline"`
	TemplateId string `json:"template_id" storm:"index"`
}
