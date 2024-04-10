package menu

import "github.com/zgwit/iot-master/v4/lib"

type Menu struct {
	Name       string   `json:"name"`
	Items      []*Item  `json:"items"`
	Domain     []string `json:"domain"` //域 admin project 或 dealer等
	Privileges []string `json:"privileges,omitempty"`
}

type Item struct {
	Name       string         `json:"name,omitempty"`
	Type       string         `json:"type,omitempty"` //internal 内部, external 嵌入web, standalone 独立弹出，默认嵌入web
	Url        string         `json:"url,omitempty"`
	Query      map[string]any `json:"query,omitempty"`
	Privileges []string       `json:"privileges,omitempty"`
}

var menus lib.Map[Menu]

func Register(name string, menu *Menu) {
	menus.Store(name, menu)
}

func Unregister(name string) {
	menus.Delete(name)
}
