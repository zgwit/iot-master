package product

import (
	"github.com/zgwit/iot-master/v4/types"
)

type Manifest struct {

	//物模型
	Properties []*Property `json:"properties,omitempty"` //属性
	Functions  []*Function `json:"functions,omitempty"`  //接口
	Events     []*Event    `json:"events,omitempty"`     //事件

	//参数
	Parameters []*Parameter `json:"parameters,omitempty"` //参数

	Validators  []*types.Validator  `json:"validators,omitempty"`
	Aggregators []*types.Aggregator `json:"aggregators,omitempty"`
}

type Parameter struct {
	Name        string     `json:"name"`
	Label       string     `json:"label"`
	Description string     `json:"description,omitempty"` //说明
	Type        types.Type `json:"type"`                  //int float ....
	Unit        string     `json:"unit"`                  //单位
	Default     any        `json:"default,omitempty"`
}

type Property struct {
	Name        string     `json:"name"`
	Label       string     `json:"label"`
	Description string     `json:"description,omitempty"` //说明
	Type        types.Type `json:"type"`                  //int float ....
	Unit        string     `json:"unit"`                  //单位
	Mode        string     `json:"mode"`                  //读取模式 r w rw
	Store       string     `json:"store"`                 // save diff
}

type Function struct {
	Name        string       `json:"name"`
	Label       string       `json:"label"`
	Description string       `json:"description,omitempty"` //说明
	Async       bool         `json:"async"`                 //异步接口
	Input       []*Parameter `json:"input"`
	Output      []*Parameter `json:"output"`
}

type Event struct {
	Name        string       `json:"name"`
	Label       string       `json:"label"`
	Description string       `json:"description,omitempty"` //说明
	Type        string       `json:"type"`                  //info alert error
	Level       uint8        `json:"level"`
	Output      []*Parameter `json:"output"`
}

//type Level uint8
