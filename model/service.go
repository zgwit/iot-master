package model

import "time"

type ServiceHistory struct {
	Id        int
	ServiceId int
	History   string
	Created   time.Time
}
