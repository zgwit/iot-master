package config

import (
	"github.com/zgwit/iot-master/args"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/iot-master/web"
	"gopkg.in/yaml.v3"
	"os"
)

//Configure 配置
type Configure struct {
	Web      web.Options      `yaml:"web"`
	Database database.Options `yaml:"database"`
	History  tsdb.Options     `yaml:"history"`
	Log      log.Options      `yaml:"log"`
}

//Config 全局配置
var Config Configure

func init() {
	Config.Web = *web.DefaultOptions()
	Config.Database = *database.DefaultOptions()
	Config.History = *tsdb.DefaultOptions()
	Config.Log = *log.DefaultOptions()

}

//Load 加载
func Load() error {
	//log.Println("加载配置")
	//从参数中读取配置文件名
	filename := args.ConfigPath

	// 如果没有文件，则使用默认信息创建
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return Store()
	} else {
		y, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer y.Close()

		d := yaml.NewDecoder(y)
		return d.Decode(&Config)
	}
	return nil
}

func Store() error {
	//log.Println("保存配置")
	//从参数中读取配置文件名
	filename := args.ConfigPath

	y, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer y.Close()

	e := yaml.NewEncoder(y)
	defer e.Close()

	return e.Encode(&Config)
}
