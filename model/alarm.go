package model

type Alarm struct {
	Id       int64  `json:"id"`
	DeviceId string `json:"device_id" xorm:"index"`
	Level    uint8  `json:"level"`
	Title    string `json:"title"`
	Message  string `json:"message,omitempty"`
	Read     bool   `json:"read,omitempty"`
	Created  Time   `json:"created,omitempty" xorm:"created"`
}

type ModParameter struct {
	Name    string  `json:"name"`
	Label   string  `json:"label"`
	Min     float64 `json:"min,omitempty"`
	Max     float64 `json:"max,omitempty"`
	Default float64 `json:"default,omitempty"`
}

type ModConstraint struct {
	Level      uint8  `json:"level"`
	Title      string `json:"title"`
	Template   string `json:"template"`
	Expression string `json:"expression"`
	Delay      uint   `json:"delay,omitempty"` //延迟时间s
	Again      uint   `json:"again,omitempty"` //再次提醒间隔s
	Total      uint   `json:"total,omitempty"` //总提醒次数
}
