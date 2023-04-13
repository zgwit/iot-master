package model

type Device struct {
	Id        string `json:"id" xorm:"pk"` //ClientID
	GatewayId string `json:"gateway_id,omitempty" xorm:"index"`
	ProductId string `json:"product_id,omitempty" xorm:"index"`

	TypeId  string `json:"type_id,omitempty" xorm:"index"`  //类型
	AreaId  string `json:"area_id,omitempty" xorm:"index"`  //区域
	GroupId string `json:"group_id,omitempty" xorm:"index"` //分组

	Name       string             `json:"name"`
	Desc       string             `json:"desc,omitempty"`
	Parameters map[string]float64 `json:"parameters,omitempty"` //模型参数，用于报警检查
	Disabled   bool               `json:"disabled,omitempty"`
	Created    Time               `json:"created,omitempty" xorm:"created"`
}

type DeviceType struct {
	Id      string `json:"id" xorm:"pk"`
	Name    string `json:"name"`
	Desc    string `json:"desc,omitempty"`
	Created Time   `json:"created,omitempty" xorm:"created"`
}

type DeviceArea struct {
	Id      string `json:"id" xorm:"pk"`
	Name    string `json:"name"`
	Desc    string `json:"desc,omitempty"`
	Created Time   `json:"created" xorm:"created"`
}

type DeviceGroup struct {
	Id      string `json:"id" xorm:"pk"`
	AreaId  string `json:"area_id" xorm:"index"`
	Name    string `json:"name"`
	Desc    string `json:"desc,omitempty"`
	Created Time   `json:"created" xorm:"created"`
}

type DeviceHistory struct {
	Id       int64  `json:"id" xorm:"pk"`
	DeviceId string `json:"device_id" xorm:"index"`
	Event    string `json:"event"`
	Created  Time   `json:"created" xorm:"created"`
}
