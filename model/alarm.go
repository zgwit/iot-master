package model

import "time"

//Alarm 数据校验
type Alarm struct {
	Disabled bool `json:"disabled"`

	//条件
	Condition string `json:"condition"`

	//重复日
	DailyChecker

	//延迟报警
	DelayChecker

	//重复报警
	RepeatChecker

	//产生告警
	AlarmContent
}

//AlarmContent 告警内容
type AlarmContent struct {
	Code    string `json:"code"`
	Level   int    `json:"level"`
	Message string `json:"message"`
}

//DeviceAlarm 设备告警
type DeviceAlarm struct {
	Id           int64 `json:"id"`
	DeviceId     int64 `json:"device_id" xorm:"index"`
	AlarmContent `xorm:"extends"`
	Created      time.Time `json:"created" xorm:"created"`
}

//ProjectAlarm 项目告警
type ProjectAlarm struct {
	Id           int64 `json:"id"`
	ProjectId    int64 `json:"project_id" xorm:"index"`
	AlarmContent `xorm:"extends"`
	Created      time.Time `json:"created" xorm:"created"`
}
