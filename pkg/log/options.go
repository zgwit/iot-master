package log

import (
	"github.com/zgwit/iot-master/v3/pkg/config"
	"github.com/zgwit/iot-master/v3/pkg/env"
)

// Options 参数
type Options struct {
	Level  string `json:"level"`
	Caller bool   `json:"caller,omitempty"`
	Text   bool   `json:"text,omitempty"`
}

func Default() Options {
	return Options{
		Level:  "trace",
		Caller: true,
		Text:   true,
	}
}

var options Options = Default()
var configure = config.AppName() + ".log.yaml"

const ENV = "IOT_MASTER_LOG_"

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
	options.Level = env.Get(ENV+"LEVEL", options.Level)
	options.Caller = env.GetBool(ENV+"CALLER", options.Caller)
	options.Text = env.GetBool(ENV+"TEXT", options.Text)
}

func (options *Options) ToEnv() []string {
	var ret []string
	if options.Level != "" {
		ret = append(ret, ENV+"LEVEL="+options.Level)
	}
	if options.Caller {
		ret = append(ret, ENV+"CALLER=TRUE")
	}
	if options.Text {
		ret = append(ret, ENV+"TEXT=TRUE")
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
