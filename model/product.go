package model

import "time"

type Product struct {
	Id      string    `json:"id" xorm:"pk"`
	ModelId string    `json:"model_id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc,omitempty"`
	Created time.Time `json:"created" xorm:"created"`
}
