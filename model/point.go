package model

//Point 数据点
type Point struct {
	Name      string        `json:"name"`
	Label     string        `json:"label"`
	Default   float64       `json:"default"`
	Type      DataType      `json:"type"`
	Precision int           `json:"precision"`
	Address   string        `json:"address"`

	LittleEndian bool `json:"little_endian"`
}
