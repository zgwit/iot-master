package model

import "time"

type Camera struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	MediaUri string `json:"media_uri"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	Deleted  time.Time `json:"-" xorm:"deleted"`
}
