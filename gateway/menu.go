package gateway

import "github.com/zgwit/iot-master/v4/menu"

func init() {
	menu.Register("gateway", &menu.Menu{
		Name:       "软网关",
		Icon:       "link",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "串口连接", Url: "serial", Type: "route"},
			{Name: "TCP客户端", Url: "client", Type: "route"},
			{Name: "TCP服务器", Url: "server", Type: "route"},
			{Name: "TCP连接", Url: "link", Type: "route"},
		},
	})
}
