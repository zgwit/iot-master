package broker

import "time"

type Gateway struct {
	Id          string `json:"id" xorm:"pk"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`

	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}
