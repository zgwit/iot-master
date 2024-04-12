package project

import "github.com/zgwit/iot-master/v4/menu"

func init() {
	menu.Register("project", &menu.Menu{
		Name:       "项目管理",
		Icon:       "apartment",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "所有项目", Url: "project", Type: "route"},
			{Name: "创建项目", Url: "project/create", Type: "route"},
		},
	})
}
