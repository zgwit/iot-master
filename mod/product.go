package mod

import "github.com/blang/semver/v4"

type Product struct {
	Id          string         `json:"id" xorm:"pk"`          //ID
	Icon        string         `json:"icon,omitempty"`        //图标
	Name        string         `json:"name"`                  //名称
	Version     semver.Version `json:"version,omitempty"`     //版本
	Type        string         `json:"type"`                  //类型：服务、应用、静态页面
	Keywords    []string       `json:"keywords,omitempty"`    //关键字
	Description string         `json:"description,omitempty"` //说明

	//物模型
	Properties []*Property `json:"properties,omitempty" xorm:"json"` //属性
	Functions  []*Function `json:"functions,omitempty" xorm:"json"`  //接口
	Events     []*Event    `json:"events,omitempty" xorm:"json"`     //事件

	//参数
	Parameters []*Parameter `json:"parameters,omitempty" xorm:"json"` //参数
}

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"` //说明
	Type        Type   `json:"type"`                  //int float ....
	Unit        string `json:"unit"`                  //单位
	Default     any    `json:"default,omitempty"`
}

type Property struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"` //说明
	Type        Type   `json:"type"`                  //int float ....
	Unit        string `json:"unit"`                  //单位
	Mode        string `json:"mode"`                  //读取模式 r w rw
}

type Function struct {
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"` //说明
	Async       bool         `json:"async"`                 //异步接口
	Input       []*Parameter `json:"input"`
	Output      []*Parameter `json:"output"`
}

type Event struct {
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"` //说明
	Type        string       `json:"type"`                  //info alert error
	Level       uint8        `json:"level"`
	Output      []*Parameter `json:"output"`
}
