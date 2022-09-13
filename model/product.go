package model

import "time"

// Product 产品
type Product struct {
	Id           string   `json:"id" xorm:"pk"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"` //厂家
	Version      string   `json:"version"`      //SEMVER
	Protocol     Protocol `json:"protocol"`

	Points  []*Point  `json:"points"`
	Pollers []*Poller `json:"pollers"`

	Created time.Time `json:"created" xorm:"created"`
}
