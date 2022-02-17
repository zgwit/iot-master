package model

import "time"

//Alarm 告警内容
type Alarm struct {
	Code    string `json:"code"`
	Level   int    `json:"level"`
	Message string `json:"message"`
}

//DeviceAlarm 设备告警
type DeviceAlarm struct {
	Alarm
	DeviceID int       `json:"device_id"`
	Created  time.Time `json:"time"`
}

//ProjectAlarm 项目告警
type ProjectAlarm struct {
	DeviceAlarm
	ProjectID int `json:"project_id"`
}
