package product

import (
	"time"
)

type Model struct {
	Id      string    `json:"id" xorm:"pk"` //ID
	Name    string    `json:"name"`         //名称
	Created time.Time `json:"created" xorm:"created"`
}
