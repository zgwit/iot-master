package model

import "time"

//User 用户
type User struct {
	ID       int       `json:"id" storm:"id,increment"`
	Username string    `json:"username" storm:"unique"`
	Nickname string    `json:"nickname,omitempty"`
	Email    string    `json:"email,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created"`
}

//Password 密码
type Password struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

//UserHistory 用户行为
type UserHistory struct {
	ID      int       `json:"id" storm:"id,increment"`
	UserID  int       `json:"user_id"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
