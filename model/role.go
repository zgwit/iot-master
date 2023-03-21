package model

type Role struct {
	Id         string   `json:"id" xorm:"pk"`
	Name       string   `json:"name"`
	Privileges []string `json:"privileges"`
	Created    Time     `json:"created,omitempty" xorm:"created"`
}

type Privilege struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`
}
