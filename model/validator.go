package model

import (
	"time"
)

type Validator struct {
	Id         string    `json:"id" xorm:"pk"`
	ProductId  string    `json:"product_id" xorm:"index"`
	Expression string    `json:"expression"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Template   string    `json:"template"`
	Level      uint      `json:"level"`
	Delay      uint      `json:"delay,omitempty"` //延迟时间s
	Again      uint      `json:"again,omitempty"` //再次提醒间隔s
	Total      uint      `json:"total,omitempty"` //总提醒次数
	Disabled   bool      `json:"disabled"`
	Created    time.Time `json:"created" xorm:"created"`
}
