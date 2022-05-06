package model

import "time"

type Pipe struct {
	Id       int       `json:"id" storm:"id,increment"`
	LinkId   int       `json:"link_id"`
	Name     string    `json:"name"`
	Addr     string    `json:"addr"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" storm:"created"`
}

type PipeEx struct {
	Pipe
	Running bool   `json:"running"`
	Link    string `json:"link"`
}
