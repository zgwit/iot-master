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
	Id        int64     `json:"id"`
	ProjectId string    `json:"project_id" xorm:"index"`
	UserId    string    `json:"user_id" xorm:"index"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created" xorm:"created"`
}

type ProjectPlugin struct {
	Id        int64     `json:"id"`
	ProjectId string    `json:"project_id" xorm:"index"`
	PluginId  string    `json:"plugin_id" xorm:"index"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created" xorm:"created"`
}

type ProjectDevice struct {
	Id        int64     `json:"id"`
	ProjectId string    `json:"project_id" xorm:"index"`
	DeviceId  string    `json:"device_id" xorm:"index"`
	Name      string    `json:"name,omitempty"` //编程别名
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created" xorm:"created"`
}
