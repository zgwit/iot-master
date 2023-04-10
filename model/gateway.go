package model

type Gateway struct {
	Id       string `json:"id" xorm:"pk"`
	Name     string `json:"name"`
	Desc     string `json:"desc,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
	Created  Time   `json:"created,omitempty" xorm:"created"`
}
