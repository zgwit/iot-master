package model

import "time"

type ComponentEvent struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type ComponentValue struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Component struct {
	Id   string `json:"id" xorm:"pk"`
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
	Events []ComponentEvent `json:"events"`

	//监听
	Values []ComponentValue `json:"values"`

	//初始化
	Create string `json:"create"`

	//写入配置
	Setup string `json:"setup"`

	//更新数据
	Update string `json:"update"`

	//产生变量 data(){return {a:1, b2}}
	Data string `json:"data"`

	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
	Deleted time.Time `json:"-" xorm:"deleted"`
}
