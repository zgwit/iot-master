package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

type Template struct {
	Id      string `json:"id" storm:"id"`
	Name    string `json:"name"`
	Version string `json:"version"` //SEMVER

	ProjectContent `storm:"inline"`

	Created time.Time `json:"created" storm:"created"`
}

type ProjectContent struct {
	Icon        string        `json:"icon"`
	Aggregators []*Aggregator `json:"aggregators"`
	Validators  []*Alarm      `json:"validators"`
	Commands    []*Command    `json:"commands"`
	Jobs        []*Job        `json:"jobs"`
	Strategies  []*Strategy   `json:"strategies"`
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

	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" storm:"created"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   int    `json:"id"`
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

type ProjectEx struct {
	Project
	Online bool
	Error  string
}
