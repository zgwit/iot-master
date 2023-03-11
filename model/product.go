package model

import "time"

type Product struct {
	Id          string          `json:"id" xorm:"pk"`
	Name        string          `json:"name"`
	Desc        string          `json:"desc,omitempty"`
	Version     string          `json:"version"`
	Properties  []ModProperty   `json:"properties" xorm:"json"`
	Functions   []ModFunction   `json:"functions" xorm:"json"`
	Events      []ModEvent      `json:"events" xorm:"json"`
	Parameters  []ModParameter  `json:"parameters"`
	Constraints []ModConstraint `json:"constraints"`

	Created time.Time `json:"created" xorm:"created"`
}

//type Model struct {
//	Values []ModProperty `json:"properties" xorm:"json"`
//	Functions  []ModFunction `json:"functions" xorm:"json"`
//	Events     []ModEvent    `json:"events" xorm:"json"`
//}

type ModProperty struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Desc  string `json:"desc,omitempty"`
	Type  string `json:"type"` //int float ....
	Unit  string `json:"unit"`
	Mode  string `json:"mode"` //r w rw
}

type ModFunction struct {
	Name   string        `json:"name"`
	Label  string        `json:"label"`
	Desc   string        `json:"desc,omitempty"`
	Async  bool          `json:"async"`
	Input  []ModArgument `json:"input"`
	Output []ModArgument `json:"output"`
}

type ModArgument struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Desc  string `json:"desc,omitempty"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
}
