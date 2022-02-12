package service

import "time"

type Link struct {
	Id        int       `json:"id" storm:"id,increment"`
	ServiceId int       `json:"service_id" storm:"index"`
	SN        string    `json:"sn"`
	Disabled  bool      `json:"disabled"`
	Created   time.Time `json:"created"`
}

type LinkHistory struct {
	Id      int       `json:"id" storm:"id,increment"`
	LinkId  int       `json:"link_id" storm:"index"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
