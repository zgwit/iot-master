package broker

import (
	"github.com/zgwit/iot-master/v4/pkg/config"
	"github.com/zgwit/iot-master/v4/pkg/env"
)

type Options struct {
	Enable bool   `json:"enable"`
	Type   string `json:"type"`
	Addr   string `json:"addr"`
}

var options Options = Default()
var configure = "broker"

const ENV = config.ENV_PREFIX + "BROKER_"

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
	options.Enable = env.GetBool(ENV+"ENABLE", options.Enable)
	options.Type = env.Get(ENV+"TYPE", options.Type)
	options.Addr = env.Get(ENV+"ADDR", options.Addr)
}

func ToEnv() []string {
	var ret []string
	if options.Enable {
		ret = append(ret, ENV+"ENABLE=TRUE")
		if options.Type != "" {
			ret = append(ret, ENV+"TYPE="+options.Type)
		}
		if options.Addr != "" {
			ret = append(ret, ENV+"ADDR="+options.Addr)
		}
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
