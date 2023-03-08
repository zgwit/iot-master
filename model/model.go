package model

type Model struct {
	Properties []Property `json:"properties" xorm:"json"`
	Functions  []Function `json:"functions" xorm:"json"`
	Events     []Event    `json:"events" xorm:"json"`
}

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
