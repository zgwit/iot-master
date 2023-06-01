package payload

type Hash map[string]any

type DevicePropertyUp struct {
	DeviceProperties
	//子设备
	Devices []DeviceProperties `json:"devices,omitempty"`
}

type Property struct {
	Name  string `json:"name"`
	Time  int64  `json:"time,omitempty"`
	Value any    `json:"value"`
}

type DeviceProperties struct {
	Id         string     `json:"id"`
	Time       int64      `json:"time,omitempty"`
	Properties []Property `json:"properties"`
}
