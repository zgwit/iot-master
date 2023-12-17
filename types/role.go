package types

import "time"

type Role struct {
	Id         string    `json:"id" xorm:"pk"`
	Name       string    `json:"name" binding:"required" xorm:"unique"`
	Privileges []string  `json:"privileges" xorm:"json"`
	Created    time.Time `json:"created,omitempty" xorm:"created"`
}

type Privilege struct {
	Id          string `json:"id" xorm:"pk"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
