package payload

import (
	"github.com/zgwit/iot-master/v3/model"
)

type Hash map[string]any

type DevicePropertyUp struct {
	DeviceProperties
	//子设备
	Devices []DeviceProperties `json:"devices,omitempty"`
}

type Property struct {
	Name      string     `json:"name"`
	Time      model.Time `json:"time,omitempty"`
	Timestamp int64      `json:"ts,omitempty"`
	Value     any        `json:"value"`
}

type DeviceProperties struct {
	Id         string     `json:"id"`
	Time       model.Time `json:"time,omitempty"`
	Timestamp  int64      `json:"ts,omitempty"`
	Properties []Property `json:"properties"`
}
