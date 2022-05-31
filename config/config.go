package config

import (
	"github.com/zgwit/iot-master/args"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

//Configure 配置
type Configure struct {
	Node     string   `yaml:"node"`
	Data     string   `yaml:"data"`
	Web      Web      `yaml:"web"`
	Database Database `yaml:"database"`
	History  History  `yaml:"history"`
	Log      Log      `yaml:"log"`
}

//Config 全局配置
var Config = Configure{
	Node:     "root",
	Data:     "data",
	Web:      WebDefault,
	Database: DatabaseDefault,
	History:  HistoryDefault,
	Log:      LogDefault,
}

func init() {
	Config.Node, _ = os.Hostname()
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
