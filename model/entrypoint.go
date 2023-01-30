package model

import (
	"time"
)

type Entrypoint struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Port int    `json:"port"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}
