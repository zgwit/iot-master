package types

import "time"

type Product struct {
	Id          string    `json:"id" xorm:"pk"`          //ID
	Icon        string    `json:"icon,omitempty"`        //图标
	Name        string    `json:"name,omitempty"`        //名称
	Url         string    `json:"url,omitempty"`         //链接
	Description string    `json:"description,omitempty"` //说明
	Keywords    []string  `json:"keywords,omitempty"`    //关键字
	Created     time.Time `json:"created" xorm:"created"`
}

type ProductVersion struct {
	ProductId string    `json:"product_id,omitempty" xorm:"PK"`
	Name      string    `json:"name,omitempty" xorm:"PK"` //版本 semver.Version
	Created   time.Time `json:"created" xorm:"created"`
}

type ProductExt struct {
	Product
}

func (p ProductExt) TableName() string {
	return "product"
}
