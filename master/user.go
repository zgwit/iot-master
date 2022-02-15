package master

import "time"

//User 用户
type User struct {
	Id       int       `json:"id" storm:"id,increment"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
}

//Password 密码
type Password struct {
	Id       int    `json:"id"  storm:"id,increment"`
	Password string `json:"password"`
}

//UserHistory 用户行为
type UserHistory struct {
	Id      int       `json:"id" storm:"id,increment"`
	UserId  int       `json:"user_id"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
