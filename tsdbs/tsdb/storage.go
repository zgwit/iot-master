package tsdb

import (
	"time"
)

type Point struct {
	Value interface{} `json:"value"`
	Time  time.Time   `json:"time"`
}

type Storage interface {
	Write(metric string, id int64, values map[string]interface{}) error
	Query(metric string, id int64, field string, start, end, window int64) ([]Point, error)
}
