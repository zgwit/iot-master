package model

import "time"

type Service struct {
	Id      string `json:"id" xorm:"pk"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Type    string `json:"type"` //tcp unix
	Address string `json:"address"`

	//TLS?
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}
