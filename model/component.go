package model

import "time"

type HmiEvent struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type HmiValue struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Component struct {
	Id   string `json:"id" storm:"id"`
	Icon string `json:"icon"` //url svg png jpg ...
	Name string `json:"name"`

	Drawer string `json:"drawer"` // "rect" | "circle" | "line" | "poly"

	//分组（默认 扩展）
	Group string `json:"group"`

	//基础配置
	Color    bool `json:"color"`    //填充色
	Stroke   bool `json:"stroke"`   //线条
	Rotation bool `json:"rotation"` //旋转
	Position bool `json:"position"` //位置

	//扩展配置项
	Properties []Hash `json:"properties"`

	//事件
	Events []HmiEvent `json:"events"`

	//监听
	Values []HmiValue `json:"values"`

	//初始化
	Create string `json:"create"`

	//写入配置
	Setup string `json:"setup"`

	//更新数据
	Update string `json:"update"`

	//产生变量 data(){return {a:1, b2}}
	Data string `json:"data"`

	Created time.Time `json:"created" storm:"created"`
}
