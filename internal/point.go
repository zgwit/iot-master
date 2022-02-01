package interval

type Point struct {
	Name    string   `json:"name"`
	Label   string   `json:"label"`
	Default float64  `json:"default"`
	Type    DataType `json:"type"`

	Code      int `json:"code"`
	Address   int `json:"address"`
	Precision int `json:"precision"`

	LittleEndian bool `json:"little_endian"`
}
