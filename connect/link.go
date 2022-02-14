package connect

import (
	"github.com/zgwit/iot-master/events"
	"time"
)

type Link interface {
	events.EventInterface

	ID() int

	Write(data []byte) error

	Close() error
}

type LinkModel struct {
	Id       int       `json:"id" storm:"id,increment"`
	TunnelId int       `json:"tunnel_id" storm:"index"`
	SN       string    `json:"sn"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created"`
}

type LinkHistory struct {
	Id      int       `json:"id" storm:"id,increment"`
	LinkId  int       `json:"link_id" storm:"index"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
