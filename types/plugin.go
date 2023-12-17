package types

import "time"

type Plugin struct {
	Id          string    `json:"id" xorm:"pk"` //ID
	Name        string    `json:"name"`         //名称
	Description string    `json:"description,omitempty"`
	Version     string    `json:"version,omitempty"`
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`
}
