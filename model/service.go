package model

import "time"

type ServiceHistory struct {
	Id        int `json:"id" storm:"id,increment"`
	ServiceId int`json:"service_id"`
	History   string`json:"history"`
	Created   time.Time`json:"created"`
}
