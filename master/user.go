package master

import "time"

type User struct {
	Id       int       `json:"id" storm:"id,increment"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
}

type Password struct {
	UserId   int    `json:"user_id"  storm:"id,increment"`
	Password string `json:"password"`
}

type UserHistory struct {
	Id      int       `json:"id" storm:"id,increment"`
	UserId  int       `json:"user_id"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
