package model

type PayloadPropertyUp struct {
	//Version string
	Id         string          `json:"id"`
	Timestamp  int64           `json:"timestamp,omitempty"`
	Properties []PayloadValue  `json:"properties,omitempty"`
	Devices    []PayloadDevice `json:"devices,omitempty"`
}

type PayloadValue struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type PayloadDevice struct {
	Id         string         `json:"id"`
	Timestamp  int64          `json:"timestamp"`
	Properties []PayloadValue `json:"properties"`
}
