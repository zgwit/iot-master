package mqtt

import (
	"github.com/zgwit/iot-master/v4/setting"
	"github.com/zgwit/iot-master/v4/types"
)

func init() {
	setting.Register(MODULE, &setting.Module{
		Name:   "MQTT连接",
		Module: MODULE,
		Title:  "MQTT连接配置",
		Form: []types.FormItem{
			{Key: "url", Label: "地址", Type: "text", Required: true, Default: ""},
			{Key: "clientid", Label: "客户端ID", Type: "text"},
			{Key: "username", Label: "用户名", Type: "text"},
			{Key: "password", Label: "密码", Type: "text"},
		},
	})
}
