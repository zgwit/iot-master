package model

import (
	"time"
)

//User 用户
type User struct {
	Id       int64     `json:"id"`
	Username string    `json:"username" xorm:"unique"`
	Nickname string    `json:"nickname,omitempty"`
	Email    string    `json:"email,omitempty"`
	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	//Deleted  time.Time `json:"-" xorm:"deleted"`
}

//Password 密码
type Password struct {
	Id       int64  `json:"id" xorm:"pk"`
	Password string `json:"password"`
}
