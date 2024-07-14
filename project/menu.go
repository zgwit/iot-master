package project

import "github.com/zgwit/iot-master/v5/menu"

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
	menu.Register("project-user", &menu.Menu{
		Name:       "用户管理",
		Icon:       "user",
		Domain:     []string{"project"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "用户", Url: "user", Type: "route"},
			{Name: "角色", Url: "role", Type: "route"},
		},
	})
}
