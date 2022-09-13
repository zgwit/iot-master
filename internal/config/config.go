package config

import (
	"github.com/zgwit/iot-master/internal/args"
	"github.com/zgwit/iot-master/internal/broker"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/pkg/log"
	"github.com/zgwit/iot-master/web"
	"gopkg.in/yaml.v3"
	"os"
)

// Configure 配置
type Configure struct {
	Node            string         `yaml:"node" json:"node"`
	Data            string         `yaml:"data" json:"data"`
	DefaultPassword string         `yaml:"default_password" json:"default_password"`
	Web             web.Options    `yaml:"web" json:"web"`
	Broker          broker.Options `yaml:"broker" json:"broker"`
	Database        db.Options     `yaml:"database" json:"database"`
	Log             log.Options    `yaml:"log" json:"log"`
}

// Config 全局配置
var Config = Configure{
	Node:            "root",
	Data:            "data",
	DefaultPassword: "123456",

	Database: db.Options{
		Type:     "sqlite",
		URL:      "sqlite3.db",
		Debug:    false,
		LogLevel: 4,
		Sync:     true,
	},

	Web: web.Options{
		Addr:     ":8080",
		Debug:    true,
		Compress: true,
	},

	Broker: broker.Options{
		Url: "internal",
	},

	Log: log.Options{
		Level:  "trace",
		Caller: true,
		Text:   false,
	},
}

func init() {
	Config.Node, _ = os.Hostname()
}

// Load 加载
func Load() error {
	//log.Println("加载配置")
	//从参数中读取配置文件名
	filename := args.ConfigPath

	// 如果没有文件，则使用默认信息创建
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return Store()
		//return nil
	} else {
		y, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer y.Close()

		d := yaml.NewDecoder(y)
		err = d.Decode(&Config)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
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

	err = e.Encode(&Config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
