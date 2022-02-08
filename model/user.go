package model

import "time"

type User struct {
	Id       int
	Username string
	Nickname string
	Email    string
	Created  time.Time
}

type Password struct {
	UserId   int
	Password string
}

type UserHistory struct {
	Id      int
	UserId  int
	History string
	Created time.Time
}
