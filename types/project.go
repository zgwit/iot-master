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
	UserId    string    `json:"user_id" xorm:"index"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}
