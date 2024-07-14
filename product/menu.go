package product

import "github.com/zgwit/iot-master/v5/menu"

func init() {
	menu.Register("product", &menu.Menu{
		Name:       "产品管理",
		Icon:       "profile",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "所有产品", Url: "product", Type: "route"},
			{Name: "创建产品", Url: "product/create", Type: "route"},
		},
	})
}
