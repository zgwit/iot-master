package model

import "time"

type Any interface{}

type Entity struct {
	Name       string              `json:"name"`
	Component  string              `json:"component"`
	Properties map[string]Any      `json:"properties"`
	Handlers   map[string][]Invoke `json:"handlers"`
	Bindings   map[string]string   `json:"bindings"`
}

type HMI struct {
	Id       string    `json:"id" storm:"id"`
	Name     string    `json:"name"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Entities []Entity  `json:"entities"`
	Created  time.Time `json:"created" storm:"created"`
}
