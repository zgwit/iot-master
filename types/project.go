package types

import "time"

type Project struct {
	Id          string    `json:"id" xorm:"pk"`
	Version     string    `json:"version,omitempty"`     //版本 semver.Version
	Icon        string    `json:"icon,omitempty"`        //图标
	Name        string    `json:"name,omitempty"`        //名称
	Url         string    `json:"url,omitempty"`         //链接
	Description string    `json:"description,omitempty"` //说明
	Keywords    []string  `json:"keywords,omitempty"`    //关键字
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`
}

type ProjectExt struct {
	Project
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
