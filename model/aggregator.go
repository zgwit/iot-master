package model

import "time"

type ModAggregator struct {
	Crontab    string `json:"crontab"`    //定时计划
	Expression string `json:"expression"` //表达式
	Type       string `json:"type"`       //聚合算法 inc dec avg count min max sum last first
	Assign     string `json:"assign"`     //赋值
}

type Aggregator struct {
	Id        string    `json:"id" xorm:"pk"`
	ProductId string    `json:"product_id" xorm:"index"`
	DeviceId  string    `json:"device_id" xorm:"index"`
	Name      string    `json:"name"`     //名称
	Desc      string    `json:"desc"`     //说明
	Disabled  bool      `json:"disabled"` //禁用
	Created   time.Time `json:"created" xorm:"created"`

	ModAggregator `xorm:"extends"`
}
