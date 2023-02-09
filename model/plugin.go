package model

import "time"

type Plugin struct {
	Id           string            `json:"id" xorm:"pk"`
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Command      string            `json:"command,omitempty"`
	Dependencies map[string]string `json:"dependencies,omitempty" xorm:"json"`
	Disabled     bool              `json:"disabled"`
	Created      time.Time         `json:"created"`
}

type License struct {
	Id      string    `json:"id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}
