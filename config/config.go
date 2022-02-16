package config

import (
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
)

//Configure 配置
type Configure struct {
	Web      *web.Option      `yaml:"web,omitempty"`
	Database *database.Option `yaml:"database,omitempty"`
	History  *tsdb.Option     `yaml:"history,omitempty"`
}

//Config 全局配置
var Config Configure = Configure{
	Web:      web.Default(),
	Database: database.Default(),
	History:  tsdb.Default(),
}

//Load 加载
func Load() error {

	return nil
}
