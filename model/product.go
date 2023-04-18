package model

type Product struct {
	Id         string         `json:"id" xorm:"pk"`
	Name       string         `json:"name"`
	Desc       string         `json:"desc,omitempty"`
	Version    string         `json:"version,omitempty"`
	Properties []ModProperty  `json:"properties,omitempty" xorm:"json"`
	Functions  []ModFunction  `json:"functions,omitempty" xorm:"json"`
	Events     []ModEvent     `json:"events,omitempty" xorm:"json"`
	Parameters []ModParameter `json:"parameters,omitempty" xorm:"json"`

	Created Time `json:"created,omitempty" xorm:"created"`
}

type ModParameter struct {
	Name    string  `json:"name"`
	Label   string  `json:"label"`
	Min     float64 `json:"min,omitempty"`
	Max     float64 `json:"max,omitempty"`
	Default float64 `json:"default,omitempty"`
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
	Mode  string `json:"mode"`  //r w rw
	Store string `json:"store"` // save diff
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

type ModEvent struct {
	Name   string        `json:"name"`
	Label  string        `json:"label"`
	Desc   string        `json:"desc,omitempty"`
	Type   string        `json:"type"` //info alert error //Level  uint8         `json:"level"`
	Output []ModArgument `json:"output"`
}
