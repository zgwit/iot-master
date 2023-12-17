package types

import "time"

type Aggregator struct {
	Crontab    string `json:"crontab"`    //定时计划
	Expression string `json:"expression"` //表达式
	Type       string `json:"type"`       //聚合算法 inc dec avg count min max sum last first
	Assign     string `json:"assign"`     //赋值
}

type ExternalAggregator struct {
	Id string `json:"id" xorm:"pk"`

	ProjectId string `json:"project_id" xorm:"index"`
	ProductId string `json:"product_id" xorm:"index"`
	DeviceId  string `json:"device_id" xorm:"index"`

	Name        string    `json:"name"`        //名称
	Description string    `json:"description"` //说明
	Disabled    bool      `json:"disabled"`    //禁用
	Created     time.Time `json:"created" xorm:"created"`

	Aggregator `xorm:"extends"`
}
