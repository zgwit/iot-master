package model

import (
	"time"
)

type ModValidator struct {
	Expression  string `json:"expression"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Level       uint   `json:"level"`
	Template    string `json:"template"`
	Delay       uint   `json:"delay,omitempty"`        //延迟时间s
	Repeat      bool   `json:"repeat,omitempty"`       //重启报警
	RepeatDelay uint   `json:"repeat_delay,omitempty"` //再次提醒间隔s
	RepeatTotal uint   `json:"repeat_total,omitempty"` //总提醒次数
}

type Validator struct {
	Id        string    `json:"id" xorm:"pk"`
	ProductId string    `json:"product_id" xorm:"index"`
	DeviceId  string    `json:"device_id" xorm:"index"`
	Disabled  bool      `json:"disabled"`
	Created   time.Time `json:"created" xorm:"created"`

	ModValidator `xorm:"extends"`
}
