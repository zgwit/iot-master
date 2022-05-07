package model

import "time"

//Point 数据点
type Point struct {
	Name      string   `json:"name"`
	Label     string   `json:"label"`
	Unit      string   `json:"unit"`
	Type      DataType `json:"type"`
	Precision int      `json:"precision"`
	Address   string   `json:"address"`
	Default   float64  `json:"default"`

	LittleEndian bool `json:"little_endian"`
}

type DataPoint struct {
	Value interface{} `json:"value"`
	Time  time.Time   `json:"time"`
}
