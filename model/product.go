package model

import "time"

type Product struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc,omitempty"`
	Model   Model     `json:"model" xorm:"JSON"`
	Created time.Time `json:"created" xorm:"created"`
}
