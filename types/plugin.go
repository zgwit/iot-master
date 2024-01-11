package types

import "time"

type Plugin struct {
	Id          string    `json:"id" xorm:"pk"`          //ID
	Version     string    `json:"version,omitempty"`     //版本 semver.Version
	Icon        string    `json:"icon,omitempty"`        //图标
	Name        string    `json:"name,omitempty"`        //名称
	Url         string    `json:"url,omitempty"`         //链接
	Description string    `json:"description,omitempty"` //说明
	Keywords    []string  `json:"keywords,omitempty"`    //关键字
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`
}

type PluginExt struct {
	Plugin
}

func (p PluginExt) TableName() string {
	return "plugin"
}
