package notify

import "time"

type Subscription struct {
	Id int64 `json:"id"`

	UserId string `json:"user_id" xorm:"index"`
	User   string `json:"user,omitempty" xorm:"<-"`

	ProductId string `json:"product_id,omitempty" xorm:"index"`
	ProjectId string `json:"project_id,omitempty" xorm:"index"`
	SpaceId   string `json:"space_id,omitempty" xorm:"index"`
	DeviceId  string `json:"device_id,omitempty" xorm:"index"`

	Product string `json:"product,omitempty" xorm:"<-"`
	Project string `json:"project,omitempty" xorm:"<-"`
	Space   string `json:"space,omitempty" xorm:"<-"`
	Device  string `json:"device,omitempty" xorm:"<-"`

	Level    uint      `json:"level"`
	Channels []string  `json:"channels" xorm:"json"`
	Disabled bool      `json:"disabled"` //禁用
	Created  time.Time `json:"created" xorm:"created"`
}
