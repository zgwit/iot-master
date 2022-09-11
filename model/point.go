package model

import (
	"time"
)

type Mode int

const (
	RW Mode = iota
	ReadOnly
	WriteOnly
)

// Point 数据点
type Point struct {
	Name    string   `json:"name"`
	Label   string   `json:"label"`
	Unit    string   `json:"unit"`
	Type    DataType `json:"type"`
	LE      bool     `json:"le"` //little endian
	Dot     int      `json:"dot"`
	Area    string   `json:"area"`
	Address string   `json:"address"`
	Store   bool     `json:"store"`
	Mode    Mode     `json:"mode"`
	//Default   float64           `json:"default"`
}

type DataPoint struct {
	Value interface{} `json:"value"`
	Time  time.Time   `json:"time"`
}
