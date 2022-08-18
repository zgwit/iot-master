package model

import "time"

type Module struct {
	Id           string            `json:"id"`
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Command      string            `json:"command,omitempty"`
	Entrypoint   string            `json:"entrypoint,omitempty"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

type License struct {
	Id      string    `json:"id"`
	Content string    `json:"content"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}
