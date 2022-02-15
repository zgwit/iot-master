package config

import (
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
)

//Configure 配置
type Configure struct {
	Web      web.Option      `yaml:"web"`
	Database database.Option `yaml:"database"`
	History  tsdb.Option     `yaml:"history"`
}

//Config 全局配置
var Config Configure = Configure{
	Web: web.Option{
		Addr: ":8080",
	},
	Database: database.Option{
		Path: ".",
	},
	History: tsdb.Option{
		DataPath: ".",
	},
}

//Load 加载
func Load() error {

	return nil
}
