package types

import "time"

type Project struct {
	Id      string `json:"id" xorm:"pk"`
	Name    string `json:"name"`
	Desc    string `json:"desc,omitempty"`
	Version string `json:"version,omitempty"`

	Created time.Time `json:"created,omitempty" xorm:"created"`
}
