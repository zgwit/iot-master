package model

import "time"

type Model struct {
	Id         string     `json:"id" xorm:"pk"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc,omitempty"`
	Version    string     `json:"version,omitempty"`
	Properties []Property `json:"properties" xorm:"json"`
	Functions  []Function `json:"functions" xorm:"json"`
	Events     []Event    `json:"events" xorm:"json"`
	Created    time.Time  `json:"created" xorm:"created"`
}

type Property struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`
	Type string `json:"type"` //int float ....
	Unit string `json:"unit"`
	Mode string `json:"mode"` //r w rw
}

type Function struct {
	Id     string     `json:"id"`
	Name   string     `json:"name"`
	Desc   string     `json:"desc,omitempty"`
	Async  bool       `json:"async"`
	Input  []Argument `json:"input"`
	Output []Argument `json:"output"`
}

type Argument struct {
	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`
	Type string `json:"type"`
	Unit string `json:"unit"`
}

type Event struct {
	Id     string     `json:"id"`
	Name   string     `json:"name"`
	Desc   string     `json:"desc,omitempty"`
	Type   string     `json:"type"` //info alert error
	Output []Argument `json:"output"`
}
