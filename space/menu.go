package space

import "github.com/zgwit/iot-master/v4/menu"

func init() {
	menu.Register("space", &menu.Menu{
		Name:       "空间管理",
		Icon:       "appstore",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "所有空间", Url: "space", Type: "route"},
			{Name: "创建空间", Url: "space/create", Type: "route"},
		},
	})
}
