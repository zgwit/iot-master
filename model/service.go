package model

import "time"

type Service struct {
	Id       int       `json:"id" storm:"id,increment"`
	Name     string    `json:"name"`
	Type     string    `json:"type"` //serial tcp-client tcp-server udp-client udp-server
	Addr     string    `json:"addr"`
	Single   bool      `json:"single"` //单线类型，针对server类型
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created"`
}

type ServiceHistory struct {
	Id        int       `json:"id" storm:"id,increment"`
	ServiceId int       `json:"service_id"`
	History   string    `json:"history"`
	Created   time.Time `json:"created"`
}
