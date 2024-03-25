package project

import (
	"github.com/zgwit/iot-master/v4/space"
	"time"
)

type ProjectUser struct {
	ProjectId string    `json:"project_id" xorm:"pk"`
	UserId    string    `json:"user_id" xorm:"pk"`
	User      string    `json:"user,omitempty" xorm:"<-"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created" xorm:"created"`
}

type Project struct {
	Id          string    `json:"id" xorm:"pk"`
	Icon        string    `json:"icon,omitempty"`        //图标
	Name        string    `json:"name,omitempty"`        //名称
	Description string    `json:"description,omitempty"` //说明
	Keywords    []string  `json:"keywords,omitempty"`    //关键字
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`

	spaces map[string]*space.Space
}
