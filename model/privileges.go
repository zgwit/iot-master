package model

var PRIVILEGES = map[string]string{
	"ALL":          "全部权限",
	"BROKER":       "总线查看",
	"BROKER-EDIT":  "总线管理",
	"PRODUCT":      "产品查看",
	"PRODUCT-EDIT": "产品管理",
	"DEVICE":       "设备查看",
	"DEVICE-EDIT":  "设备管理",
	"ALARM":        "报警查看",
	"ALARM-EDIT":   "报警管理",
	"USER":         "用户查看",
	"USER-EDIT":    "用户管理",
	"PLUGIN":       "插件查看",
	"PLUGIN-EDIT":  "插件管理",
	"SETTING":      "系统设置",
}
