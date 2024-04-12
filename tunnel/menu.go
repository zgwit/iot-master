package tunnel

import "github.com/zgwit/iot-master/v4/menu"

func init() {
	menu.Register("tunnel", &menu.Menu{
		Name:       "连接管理",
		Icon:       "link",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "TCP服务器", Url: "server", Type: "route"},
			{Name: "TCP连接", Url: "link", Type: "route"},
			{Name: "TCP客户端", Url: "client", Type: "route"},
			{Name: "串口连接", Url: "serial", Type: "route"},
		},
	})
}
