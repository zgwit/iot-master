package model

import "time"

//Link 链接模型
type Link struct {
	Id       int64     `json:"id"`
	TunnelId int64     `json:"tunnel_id" xorm:"index"`
	SN       string    `json:"sn" xorm:"index 'sn'"`
	Name     string    `json:"name"`
	Remote   string    `json:"remote"`
	Disabled bool      `json:"disabled"`
	Last     time.Time `json:"last"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	Deleted  time.Time `json:"-" xorm:"deleted"`
	//Protocol *Protocol `json:"protocol"`
}

type LinkEx struct {
	Link `xorm:"extends"`
	Running bool   `json:"running"`
	Tunnel  string `json:"tunnel"`
}

func (link *LinkEx) TableName() string {
	return "link"
}
