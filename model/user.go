package model

import (
	"time"
)

// User 用户
type User struct {
	Id       int64     `json:"id"`
	Username string    `json:"username" xorm:"unique"`
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}

// Password 密码
type Password struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

type UserHistory struct {
	Id      int64     `json:"id"`
	UserId  int64     `json:"user_id"`
	Event   string    `json:"event"`
	Created time.Time `json:"created" xorm:"created"`
}
