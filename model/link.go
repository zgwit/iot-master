package model

import "time"

type LinkHistory struct {
	Id      int
	LinkId  int
	History string
	Created time.Time
}
