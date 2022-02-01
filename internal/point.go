package interval

type Point struct {
	Name    string   `json:"name"`
	Label   string   `json:"label"`
	Default float64  `json:"default"`
	Type    DataType `json:"type"`

	Code      int `json:"code"`
	Address   int `json:"address"`
	//TODO address2
	Precision int `json:"precision"`

	LittleEndian bool `json:"little_endian"`
}
