package model

type TopicPropertyUp struct {
	//Version string
	Id         string        `json:"id"`
	Timestamp  int64         `json:"timestamp,omitempty"`
	Properties []TopicValue  `json:"properties,omitempty"`
	Devices    []TopicDevice `json:"devices,omitempty"`
}

type TopicValue struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type TopicDevice struct {
	Id         string       `json:"id"`
	Timestamp  int64        `json:"timestamp"`
	Properties []TopicValue `json:"properties"`
}
