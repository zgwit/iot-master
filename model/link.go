package model

import "time"

//Link 链接模型
type Link struct {
	ID       int       `json:"id" storm:"id,increment"`
	TunnelID int       `json:"tunnel_id" storm:"index"`
	SN       string    `json:"sn"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created"`
}

//LinkHistory 链接历史
type LinkHistory struct {
	ID      int       `json:"id" storm:"id,increment"`
	LinkID  int       `json:"link_id" storm:"index"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
