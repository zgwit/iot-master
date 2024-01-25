package types

import "time"

type Space struct {
	Id          string    `json:"id" xorm:"pk"`
	Name        string    `json:"name,omitempty"`        //名称
	Description string    `json:"description,omitempty"` //说明
	ProjectId   string    `json:"project_id,omitempty" xorm:"index"`
	Project     string    `json:"project,omitempty" xorm:"<-"`
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`
}

type SpaceDevice struct {
	SpaceId  string    `json:"space_id" xorm:"pk"`
	DeviceId string    `json:"device_id" xorm:"pk"`
	Name     string    `json:"name,omitempty"` //编程别名
	Created  time.Time `json:"created" xorm:"created"`
}
