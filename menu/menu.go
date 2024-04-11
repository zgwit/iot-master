package menu

import "github.com/zgwit/iot-master/v4/lib"

type Menu struct {
	Name       string   `json:"name"`
	Domain     []string `json:"domain"` //域 admin project 或 dealer等
	Privileges []string `json:"privileges,omitempty"`
	Items      []*Item  `json:"items"`
}

type Item struct {
	Name       string         `json:"name,omitempty"`
	Type       string         `json:"type,omitempty"` //route 路由, web 嵌入web, window 独立弹出
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
