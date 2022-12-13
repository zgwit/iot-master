package model

import "time"

type Event struct {
	Type string `json:"type"`
	Data string `json:"data,omitempty"`
}

type UserEvent struct {
	Id      int64 `json:"id"`
	UserId  int64 `json:"user_id"`
	Event   `xorm:"extends"`
	Created time.Time `json:"created" xorm:"created"`
}

type GatewayEvent struct {
	Id        int64  `json:"id"`
	GatewayId string `json:"gateway_id"`
	Event     `xorm:"extends"`
	Created   time.Time `xorm:"created"`
}

type TunnelEvent struct {
	Id       int64  `json:"id"`
	TunnelId string `json:"tunnel_id"`
	Event    `xorm:"extends"`
	Created  time.Time `xorm:"created"`
}

type ServerEvent struct {
	Id       int64  `json:"id"`
	ServerId string `json:"server_id"`
	Event    `xorm:"extends"`
	Created  time.Time `xorm:"created"`
}

type DeviceEvent struct {
	Id       int64  `json:"id"`
	DeviceId string `json:"device_id"`
	Event    `xorm:"extends"`
	Created  time.Time `xorm:"created"`
}

type ProjectEvent struct {
	Id        int64  `json:"id"`
	ProjectId string `json:"project_id"`
	Event     `xorm:"extends"`
	Created   time.Time `xorm:"created"`
}
