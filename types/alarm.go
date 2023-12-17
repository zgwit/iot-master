package types

import (
	"time"
)

type Alarm struct {
	Id int64 `json:"id"`

	ProjectId string `json:"project_id,omitempty" xorm:"index"`
	ProductId string `json:"product_id,omitempty" xorm:"index"`
	DeviceId  string `json:"device_id,omitempty" xorm:"index"`

	Type    string    `json:"type,omitempty"`
	Title   string    `json:"title"`
	Message string    `json:"message,omitempty"`
	Level   uint      `json:"level"`
	Read    bool      `json:"read,omitempty"`
	Created time.Time `json:"created,omitempty" xorm:"created"`
}

type AlarmEx struct {
	Alarm   `xorm:"extends"`
	Project string `json:"project,omitempty" xorm:"<-"`
	Product string `json:"product,omitempty" xorm:"<-"`
	Device  string `json:"device,omitempty" xorm:"<-"`
}

func (a *AlarmEx) TableName() string {
	return "alarm"
}

type Subscription struct {
	Id     int64  `json:"id"`
	UserId string `json:"user_id" xorm:"index"`

	ProjectId string `json:"project_id,omitempty" xorm:"index"`
	ProductId string `json:"product_id,omitempty" xorm:"index"`
	DeviceId  string `json:"device_id,omitempty" xorm:"index"`
	Project   string `json:"project,omitempty" xorm:"<-"`
	Product   string `json:"product,omitempty" xorm:"<-"`
	Device    string `json:"device,omitempty" xorm:"<-"`

	Level    uint      `json:"level"`
	Channels []string  `json:"channels" xorm:"json"`
	Disabled bool      `json:"disabled"` //禁用
	Created  time.Time `json:"created" xorm:"created"`
}

// 通知
type Notification struct {
	Id       int64     `json:"id,omitempty"`
	AlarmId  int64     `json:"alarm_id,omitempty" xorm:"index"`
	UserId   string    `json:"user_id,omitempty" xorm:"index"`
	Channels []string  `json:"channels" xorm:"json"`
	Created  time.Time `json:"created" xorm:"created"`
}
