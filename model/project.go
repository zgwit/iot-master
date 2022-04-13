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

//ProjectEvent 项目历史
type ProjectEvent struct {
	ID        int       `json:"id" storm:"id,increment"`
	ProjectID int       `json:"project_id"`
	Event     string    `json:"event"`
	Created   time.Time `json:"created" storm:"created"`
}
