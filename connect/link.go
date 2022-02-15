package connect

import (
	"github.com/zgwit/iot-master/events"
	"time"
)

//Link 链接
type Link interface {
	events.EventInterface

	ID() int

	Write(data []byte) error

	Close() error
}

//LinkModel 链接模型
type LinkModel struct {
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
