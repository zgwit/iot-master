package model

import (
	"time"
)

type Alarm struct {
	Id          int64     `json:"id"`
	ProductId   string    `json:"product_id,omitempty" xorm:"index"`
	DeviceId    string    `json:"device_id,omitempty" xorm:"index"`
	ValidatorId string    `json:"validator_id,omitempty" xorm:"index"`
	Product     string    `json:"product,omitempty" xorm:"-"`
	Device      string    `json:"device,omitempty" xorm:"-"`
	Type        string    `json:"type,omitempty"`
	Title       string    `json:"title"`
	Message     string    `json:"message,omitempty"`
	Level       uint      `json:"level"`
	Read        bool      `json:"read,omitempty"`
	Created     time.Time `json:"created,omitempty" xorm:"created"`
}

type Subscription struct {
	Id          int64     `json:"id"`
	UserId      string    `json:"user_id" xorm:"index"`
	ProductId   string    `json:"product_id,omitempty" xorm:"index"`
	DeviceId    string    `json:"device_id,omitempty" xorm:"index"`
	ValidatorId string    `json:"validator_id,omitempty" xorm:"index"`
	Product     string    `json:"product,omitempty" xorm:"-"`
	Device      string    `json:"device,omitempty" xorm:"-"`
	Level       uint      `json:"level"`
	Channels    []string  `json:"channels" xorm:"json"`
	Sms         bool      `json:"sms,omitempty"`
	Voice       bool      `json:"voice,omitempty"`
	Disabled    bool      `json:"disabled"` //禁用
	Created     time.Time `json:"created" xorm:"created"`
}

// 通知
type Notification struct {
	Id       int64     `json:"id,omitempty"`
	AlarmId  string    `json:"alarm_id,omitempty" xorm:"index"`
	UserId   string    `json:"user_id,omitempty" xorm:"index"`
	Channels []string  `json:"channels" xorm:"json"`
	Created  time.Time `json:"created" xorm:"created"`
}
