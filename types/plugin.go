package types

import "time"

type Plugin struct {
	Id       string    `json:"id" xorm:"pk"` //ID
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" xorm:"created"`
}

type PluginExt struct {
	Plugin
	ManifestBase
}

func (p PluginExt) TableName() string {
	return "plugin"
}
