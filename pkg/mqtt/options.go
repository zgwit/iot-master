package mqtt

import (
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/pkg/env"
)

type Options struct {
	Url      string `json:"url,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var options Options = Default()
var configure = "mqtt"

const ENV = config.ENV_PREFIX + "MQTT_"

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
	options.Url = env.Get(ENV+"URL", options.Url)
	options.ClientId = env.Get(ENV+"CLIENT_ID", options.ClientId)
	options.Username = env.Get(ENV+"USERNAME", options.Username)
	options.Password = env.Get(ENV+"PASSWORD", options.Password)
}

func (options *Options) ToEnv() []string {
	return []string{
		ENV + "URL=" + options.Url,
		ENV + "CLIENT_ID=" + options.ClientId,
		ENV + "USERNAME=" + options.Username,
		ENV + "PASSWORD=" + options.Password}
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
