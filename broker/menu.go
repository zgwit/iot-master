package broker

import "github.com/zgwit/iot-master/v4/menu"

func init() {
	menu.Register("broker", &menu.Menu{
		Name:       "MQTT服务器",
		Icon:       "cluster",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "服务器", Url: "broker", Type: "route"},
			{Name: "网关", Url: "gateway", Type: "route"},
		},
	})
}
