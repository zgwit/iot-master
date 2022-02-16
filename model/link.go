package model

import "time"

//Link 链接模型
type Link struct {
	Id       int       `json:"id" storm:"id,increment"`
	TunnelId int       `json:"tunnel_id" storm:"index"`
	SN       string    `json:"sn"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created"`
}

//LinkHistory 链接历史
type LinkHistory struct {
	Id      int       `json:"id" storm:"id,increment"`
	LinkId  int       `json:"link_id" storm:"index"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
