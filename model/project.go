package model

import (
	"time"
)

type Template struct {
	Id      string `json:"id" xorm:"pk"`
	Name    string `json:"name"`
	Version string `json:"version"` //SEMVER

	Products []*TemplateProduct `json:"products"`

	ProjectContent `xorm:"extends"`

	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
	//Deleted time.Time `json:"-" xorm:"deleted"`
}

type ProjectContent struct {
	Hmi         string        `json:"hmi"`
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

	//Context map[string]interface{} `json:"context"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	//Deleted  time.Time `json:"-" xorm:"deleted"`
}

type ProjectEx struct {
	Project  `xorm:"extends"`
	Running  bool   `json:"running"`
	Template string `json:"template"`
}

func (p *ProjectEx) TableName() string {
	return "project"
}

//ProjectDevice 项目的设备
type ProjectDevice struct {
	Id   int64  `json:"id"`
	Name string `json:"name"` //编程名
}

type TemplateProduct struct {
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
