package device

import "github.com/zgwit/iot-master/v4/menu"

func init() {
	menu.Register("device", &menu.Menu{
		Name:       "设备管理",
		Icon:       "block",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "所有设备", Url: "device", Type: "route"},
			{Name: "创建设备", Url: "device/create", Type: "route"},
		},
	})
}
