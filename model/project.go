package model

import (
	"github.com/zgwit/iot-master/calc"
	"time"
)

type Template struct {
	Id      string `json:"id" xorm:"pk"`
	Name    string `json:"name"`
	Version string `json:"version"` //SEMVER

	Elements []*TemplateElement `json:"elements"`

	ProjectContent `xorm:"extends"`

	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
	Deleted time.Time `json:"-" xorm:"deleted"`
}

type ProjectContent struct {
	HMI         string        `json:"hmi" xorm:"'hmi'"`
	Aggregators []*Aggregator `json:"aggregators"`
	Jobs        []*Job        `json:"jobs"`
	Alarms      []*Alarm      `json:"alarms"`
	Strategies  []*Strategy   `json:"strategies"`
}

//Project 项目
type Project struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Devices []*ProjectDevice `json:"devices"`
	//Devices []int `json:"devices"`

	TemplateId     string `json:"template_id,omitempty"`
	ProjectContent `xorm:"extends"`

	Context calc.Context `json:"context"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	Deleted  time.Time `json:"-" xorm:"deleted"`
}

type ProjectEx struct {
	Project  `xorm:"extends"`
	Running  bool   `json:"running"`
	Template string `json:"template"`
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   int64  `json:"id"`
	Name string `json:"name"` //编程名
}

type TemplateElement struct {
	Id   string `json:"id"`
	Name string `json:"name"` //编程名
}

type ProjectHistory struct {
	Project   `xorm:"extends"`
	ProjectId int64 `json:"project_id" xorm:"index"`
}

type TemplateHistory struct {
	Template   `xorm:"extends"`
	TemplateId string `json:"template_id" xorm:"index"`
}
