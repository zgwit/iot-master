package model

import (
	"encoding/gob"
	"time"
)

//User 用户
type User struct {
	Id       int       `json:"id" storm:"id,increment"`
	Username string    `json:"username" storm:"unique"`
	Nickname string    `json:"nickname,omitempty"`
	Email    string    `json:"email,omitempty"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" storm:"created"`
}

func init() {
	gob.Register(User{})
}

//Password 密码
type Password struct {
	Id       int    `json:"id" storm:"id"`
	Password string `json:"password"`
}

//UserEvent 用户行为
type UserEvent struct {
	Id      int       `json:"id" storm:"id,increment"`
	UserId  int       `json:"user_id"`
	Event   string    `json:"event"`
	Created time.Time `json:"created" storm:"created"`
}
