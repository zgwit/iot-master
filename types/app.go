package types

import "time"

type AppEntry struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"`
}

type App struct {
	Id      string      `json:"id" xorm:"pk"`
	Name    string      `json:"name"`
	Icon    string      `json:"icon,omitempty"`
	Entries []*AppEntry `json:"entries" xorm:"json"`
	Desc    string      `json:"desc,omitempty"`
	Type    string      `json:"type"` //tcp unix
	Address string      `json:"address"`
	Auth    string      `json:"auth"`             //鉴权机制，继承 无 自定义，inherit none custom
	Hidden  bool        `json:"hidden,omitempty"` //隐藏，适用于服务
}

type AppHistory struct {
	Id      int64     `json:"id"`
	AppId   string    `json:"app_id"`
	Event   string    `json:"event"`
	Created time.Time `json:"created" xorm:"created"`
}
