package model

import "time"

type Product struct {
	Id         string     `json:"id" xorm:"pk"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc,omitempty"`
	Version    string     `json:"version"`
	Properties []Property `json:"properties" xorm:"json"`
	Functions  []Function `json:"functions" xorm:"json"`
	Events     []Event    `json:"events" xorm:"json"`
	Created    time.Time  `json:"created" xorm:"created"`
}

//type Model struct {
//	Properties []Property `json:"properties" xorm:"json"`
//	Functions  []Function `json:"functions" xorm:"json"`
//	Events     []Event    `json:"events" xorm:"json"`
//}

type Property struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Desc  string `json:"desc,omitempty"`
	Type  string `json:"type"` //int float ....
	Unit  string `json:"unit"`
	Mode  string `json:"mode"` //r w rw
}

type Function struct {
	Name   string     `json:"name"`
	Label  string     `json:"label"`
	Desc   string     `json:"desc,omitempty"`
	Async  bool       `json:"async"`
	Input  []Argument `json:"input"`
	Output []Argument `json:"output"`
}

type Argument struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Desc  string `json:"desc,omitempty"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
}

type Event struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Desc  string `json:"desc,omitempty"`
	//Type   string     `json:"type"` //info alert error
	Level  uint8      `json:"level"`
	Output []Argument `json:"output"`
}
