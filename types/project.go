package types

import "time"

type Project struct {
	Id       string    `json:"id" xorm:"pk"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" xorm:"created"`
}

type ProjectExt struct {
	Project
	ManifestBase
}

func (p ProjectExt) TableName() string {
	return "project"
}

type ProjectUser struct {
	ProjectId string    `json:"project_id" xorm:"pk"`
	UserId    string    `json:"user_id" xorm:"pk"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created" xorm:"created"`
}

type ProjectPlugin struct {
	ProjectId string    `json:"project_id" xorm:"pk"`
	PluginId  string    `json:"plugin_id" xorm:"pk"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created" xorm:"created"`
}

type ProjectDevice struct {
	ProjectId string    `json:"project_id" xorm:"pk"`
	DeviceId  string    `json:"device_id" xorm:"pk"`
	Name      string    `json:"name,omitempty"` //编程别名
	Created   time.Time `json:"created" xorm:"created"`
}
