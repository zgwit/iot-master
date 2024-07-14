package space

import "github.com/zgwit/iot-master/v5/menu"

func init() {
	menu.Register("space", &menu.Menu{
		Name:       "空间管理",
		Icon:       "appstore",
		Domain:     []string{"project"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "所有空间", Url: "space", Type: "route"},
			{Name: "创建空间", Url: "space/create", Type: "route"},
		},
	})
}
