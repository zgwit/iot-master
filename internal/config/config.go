package config

import (
	"github.com/zgwit/iot-master/v3/internal/args"
	"github.com/zgwit/iot-master/v3/internal/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"gopkg.in/yaml.v3"
	"os"
)

// Configure 配置
type Configure struct {
	Data     string      `yaml:"data" json:"data"`
	Web      uint16      `yaml:"web" json:"web"`
	Oem      OEM         `yaml:"oem" json:"oem"`
	Database db.Options  `yaml:"database" json:"database"`
	Log      log.Options `yaml:"log" json:"log"`
}

// Config 全局配置
var Config = Configure{
	Data: "data",
	Web:  8888,
	Oem: OEM{
		Title:     "物联大师",
		Logo:      "",
		Company:   "无锡真格智能科技有限公司",
		Copyright: "©2023",
	},
	Database: db.Options{
		Type:     "mysql",
		URL:      "root:123456@localhost:3306/master?charset=utf8",
		Debug:    false,
		LogLevel: 4,
		Sync:     true,
	},

	Log: log.Options{
		Level:  "trace",
		Caller: true,
		Text:   true,
	},
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
