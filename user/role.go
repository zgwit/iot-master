package user

import (
	"github.com/zgwit/iot-master/v4/db"
	"time"
)

func init() {
	db.Register(new(Role))
}

type Role struct {
	Id          string    `json:"id" xorm:"pk"`
	Name        string    `json:"name,omitempty"`        //名称
	Description string    `json:"description,omitempty"` //说明
	Privileges  []string  `json:"privileges,omitempty"`
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`
}
