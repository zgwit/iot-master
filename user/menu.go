package user

import "github.com/zgwit/iot-master/v5/menu"

func init() {
	menu.Register("user", &menu.Menu{
		Name:       "用户管理",
		Icon:       "user",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "所有用户", Url: "user", Type: "route"},
			{Name: "所有角色", Url: "role", Type: "route"},
		},
	})
}
