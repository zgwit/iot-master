package notify

import "time"

type Notification struct {
	Id int64 `json:"id,omitempty"`

	AlarmId int64  `json:"alarm_id,omitempty" xorm:"index"`
	Title   string `json:"title,omitempty" xorm:"<-"`

	UserId string `json:"user_id,omitempty" xorm:"index"`
	User   string `json:"user,omitempty" xorm:"<-"`

	Channels []string  `json:"channels" xorm:"json"`
	Created  time.Time `json:"created" xorm:"created"`
}
