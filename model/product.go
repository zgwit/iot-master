package model

import "time"

type Product struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc,omitempty"`
	ModelId string    `json:"model_id"`
	Created time.Time `json:"created" xorm:"created"`
}
