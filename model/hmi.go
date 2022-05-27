package model

import "time"

type Entity struct {
	Name       string              `json:"name"`
	Component  string              `json:"component"`
	Properties map[string]Any      `json:"properties"`
	Handlers   map[string][]Invoke `json:"handlers"`
	Bindings   map[string]string   `json:"bindings"`
}

type Hmi struct {
	Id       string    `json:"id" xorm:"pk"`
	Name     string    `json:"name"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Snap     string    `json:"snap"`
	Entities []Entity  `json:"entities"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	Deleted  time.Time `json:"-" xorm:"deleted"`
}
