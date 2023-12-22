package types

import "time"

type Product struct {
	Id       string    `json:"id" xorm:"pk"` //ID
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" xorm:"created"`
}

type ProductExt struct {
	Product
	ManifestBase
}

func (p ProductExt) TableName() string {
	return "product"
}
