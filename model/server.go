package model

import (
	"time"
)

type Server struct {
	Id        string          `json:"id" xorm:"pk"`
	GatewayId string          `json:"gateway_id"`
	Name      string          `json:"name"`
	Type      string          `json:"type"` //tcp udp
	Addr      string          `json:"addr"`
	Protocol  Protocol        `json:"protocol" xorm:"JSON"`
	Devices   []DefaultDevice `json:"devices" xorm:"JSON"` //默认设备
	Disabled  bool            `json:"disabled"`
	Created   time.Time       `json:"created" xorm:"created"`
}
