package pool

import (
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/pkg/env"
	"strconv"
)

// Options 参数
type Options struct {
	Size int `json:"size"`
}

func Default() Options {
	return Options{
		Size: 10000,
	}
}

var options Options = Default()
var configure = "pool"

const ENV = config.ENV_PREFIX + "POOL_"

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
	options.Size = env.GetInt(ENV+"SIZE", options.Size)
}

func (options *Options) ToEnv() []string {
	var ret []string
	if options.Size > 0 {
		ret = append(ret, ENV+"Size="+strconv.Itoa(options.Size))
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
