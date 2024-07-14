package gateway

import "github.com/zgwit/iot-master/v5/menu"

func init() {
	menu.Register("broker", &menu.Menu{
		Name:       "MQTT服务器",
		Icon:       "cluster",
		Domain:     []string{"admin"},
		Privileges: nil,
		Items: []*menu.Item{
			{Name: "网关", Url: "gateway", Type: "route"},
		},
	})
}
