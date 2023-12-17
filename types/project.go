package types

import "time"

type Project struct {
	Id          string    `json:"id" xorm:"pk"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Version     string    `json:"version,omitempty"`
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created,omitempty" xorm:"created"`
}

type ProjectUser struct {
	Id        string    `json:"id" xorm:"pk"`
	ProjectId string    `json:"project_id" xorm:"index"`
	Username  string    `json:"username"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Roles     []string  `json:"roles,omitempty" xorm:"json"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}

type ProjectUserPassword struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}
