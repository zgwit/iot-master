package model

import "time"

type LinkHistory struct {
	Id      int       `json:"id" storm:"id,increment"`
	LinkId  int       `json:"link_id"`
	History string    `json:"history"`
	Created time.Time `json:"created"`
}
