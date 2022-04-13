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
	ID       int `json:"id" storm:"id,increment"`
	DeviceID int `json:"device_id" storm:"index"`
	Alarm    `storm:"inline"`
	Created  time.Time `json:"created"`
}

//ProjectAlarm 项目告警
type ProjectAlarm struct {
	ID        int `json:"id" storm:"id,increment"`
	ProjectID int `json:"project_id" storm:"index"`
	Alarm     `storm:"inline"`
	Created   time.Time `json:"created"`
}
