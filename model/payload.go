package model

type UpPropertyPayload struct {
	PayloadDevice
	//子设备
	Devices []PayloadDevice `json:"devices,omitempty"`
}

type ValuePayload struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Value     any    `json:"value"`
}

type PayloadDevice struct {
	Id         string         `json:"id"`
	Timestamp  int64          `json:"timestamp,omitempty"`
	Properties []ValuePayload `json:"properties"`
}
