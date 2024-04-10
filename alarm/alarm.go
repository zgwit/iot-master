package alarm

import (
	"github.com/zgwit/iot-master/v4/db"
	"time"
)

func init() {
	db.Register(new(Alarm))
}

type Alarm struct {
	Id int64 `json:"id"`

	ProductId string `json:"product_id,omitempty" xorm:"index"`
	ProjectId string `json:"project_id,omitempty" xorm:"index"`
	SpaceId   string `json:"space_id,omitempty" xorm:"index"`
	DeviceId  string `json:"device_id,omitempty" xorm:"index"`

	Product string `json:"product,omitempty" xorm:"<-"`
	Project string `json:"project,omitempty" xorm:"<-"`
	Space   string `json:"space,omitempty" xorm:"<-"`
	Device  string `json:"device,omitempty" xorm:"<-"`

	Level   uint   `json:"level,omitempty"`
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`

	Read    bool      `json:"read,omitempty"`
	Created time.Time `json:"created,omitempty" xorm:"created"`
}
