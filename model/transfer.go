package model

import "time"

type Transfer struct {
	Id       int64     `json:"id"`
	TunnelId int64     `json:"tunnel_id"`
	Name     string    `json:"name"`
	Port     int       `json:"port"`
	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	//Deleted  time.Time `json:"-" xorm:"deleted"`
}

type TransferEx struct {
	Transfer `xorm:"extends"`
	Running  bool   `json:"running"`
	Tunnel   string `json:"tunnel"`
}

func (p *TransferEx) TableName() string {
	return "transfer"
}
