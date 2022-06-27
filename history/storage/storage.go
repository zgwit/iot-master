package storage

import (
	"iot-master/helper"
	"time"
)

type Point struct {
	Value interface{} `json:"value"`
	Time  time.Time   `json:"time"`
}

type Storage interface {
	Open(options helper.Options) error
	Close() error
	Write(id int64, values map[string]interface{}) error
	Query(id int64, field string, start, end, window string) ([]Point, error)
}
