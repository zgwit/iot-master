package types

import "time"

type Device struct {
	Id        string `json:"id" xorm:"pk"` //ClientID
	GatewayId string `json:"gateway_id,omitempty" xorm:"index"`
	ProductId string `json:"product_id,omitempty" xorm:"index"`
	Gateway   string `json:"gateway,omitempty" xorm:"<-"`
	Product   string `json:"product,omitempty" xorm:"<-"`

	Name       string             `json:"name"`
	Desc       string             `json:"desc,omitempty"`
	Parameters map[string]float64 `json:"parameters,omitempty" xorm:"json"` //模型参数，用于报警检查
	Disabled   bool               `json:"disabled,omitempty"`
	Created    time.Time          `json:"created,omitempty" xorm:"created"`

	Online bool `json:"online,omitempty" xorm:"-"`
}

type DeviceHistory struct {
	Id       int64     `json:"id" xorm:"pk"`
	DeviceId string    `json:"device_id" xorm:"index"`
	Event    string    `json:"event"`
	Created  time.Time `json:"created" xorm:"created"`
}
