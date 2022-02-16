package config

import (
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
)

//Configure 配置
type Configure struct {
	Web      *web.Options      `yaml:"web,omitempty"`
	Database *database.Options `yaml:"database,omitempty"`
	History  *tsdb.Options     `yaml:"history,omitempty"`
}

//Config 全局配置
var Config Configure

//Load 加载
func Load() error {

	return nil
}
