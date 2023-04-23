package web

import (
	"github.com/zgwit/iot-master/v3/pkg/config"
	"github.com/zgwit/iot-master/v3/pkg/env"
)

// Options 参数
type Options struct {
	Addr  string `yaml:"addr" json:"addr"`
	Debug bool   `yaml:"debug,omitempty" json:"debug,omitempty"`
	Cors  bool   `json:"cors,omitempty" json:"cors,omitempty"`
	Gzip  bool   `json:"gzip,omitempty" json:"gzip,omitempty"`
}

func Default() Options {
	return Options{
		Addr:  ":8080",
		Debug: true,
		Cors:  false,
		Gzip:  true,
	}
}

var options Options = Default()
var configure = config.AppName() + ".web.yaml"

const ENV = "IOT_MASTER_WEB_"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	//首先加载环境变量
	FromEnv()
}

func FromEnv() {
	options.Addr = env.Get(ENV+"ADDR", options.Addr)
	options.Debug = env.GetBool(ENV+"DEBUG", options.Debug)
	options.Cors = env.GetBool(ENV+"CORS", options.Cors)
	options.Gzip = env.GetBool(ENV+"GZIP", options.Gzip)
}

func ToEnv() []string {
	ret := []string{ENV + "ADDR=" + options.Addr}
	if options.Debug {
		ret = append(ret, ENV+"DEBUG=TRUE")
	}
	if options.Cors {
		ret = append(ret, ENV+"CORS=TRUE")
	}
	if options.Gzip {
		ret = append(ret, ENV+"GZIP=TRUE")
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
