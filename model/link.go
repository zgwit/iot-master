package model

import "time"

//Link 链接模型
type Link struct {
	ID       int       `json:"id" storm:"id,increment"`
	TunnelID int       `json:"tunnel_id" storm:"index"`
	SN       string    `json:"sn"`
	Disabled bool      `json:"disabled"`
	//Protocol *Protocol `json:"protocol"`
	Created  time.Time `json:"created" storm:"created"`
}

//LinkEvent 链接历史
type LinkEvent struct {
	ID      int       `json:"id" storm:"id,increment"`
	LinkID  int       `json:"link_id" storm:"index"`
	Event   string    `json:"event"`
	Created time.Time `json:"created" storm:"created"`
}
