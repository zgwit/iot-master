package master

import "github.com/zgwit/iot-master/model"

//Point 数据点
type Point struct {
	Name    string         `json:"name"`
	Label   string         `json:"label"`
	Default float64        `json:"default"`
	Type    model.DataType `json:"type"`

	Code    int `json:"code"`
	Address int `json:"address"`
	//TODO address2
	Precision int `json:"precision"`

	LittleEndian bool `json:"little_endian"`
}
