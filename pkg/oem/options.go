package oem

import (
	"github.com/zgwit/iot-master/v4/pkg/config"
	"github.com/zgwit/iot-master/v4/pkg/env"
)

type Options struct {
	Title     string `yaml:"title,omitempty" json:"title,omitempty"`
	Logo      string `yaml:"logo,omitempty" json:"logo,omitempty"`
	Company   string `yaml:"company,omitempty" json:"company,omitempty"`
	Copyright string `yaml:"copyright,omitempty" json:"copyright,omitempty"`
}

func Default() Options {
	return Options{
		Title:     "物联大师",
		Logo:      "",
		Company:   "无锡真格智能科技有限公司",
		Copyright: "©2023",
	}
}

var options Options = Default()
var configure = config.AppName() + ".oem.yaml"

const ENV = "IOT_MASTER_OEM_"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	//首先加载环境变量
	options.FromEnv()
}

func (options *Options) FromEnv() {
	options.Title = env.Get(ENV+"ADDR", options.Title)
	options.Logo = env.Get(ENV+"DEBUG", options.Logo)
	options.Company = env.Get(ENV+"CORS", options.Company)
	options.Copyright = env.Get(ENV+"GZIP", options.Copyright)
}

func (options *Options) ToEnv() []string {
	return []string{
		ENV + "ADDR=" + options.Title,
		ENV + "DEBUG=" + options.Logo,
		ENV + "CORS=" + options.Company,
		ENV + "GZIP=" + options.Copyright}
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
