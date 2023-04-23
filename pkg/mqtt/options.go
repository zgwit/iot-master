package mqtt

import (
	"encoding/json"
	"github.com/zgwit/iot-master/v3/pkg/config"
	"os"
)

type Options struct {
	Url      string `json:"url,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var options Options = Default()
var configure = config.AppName() + ".mqtt.yaml"

const ENV = "iot-master-options-mqtt"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	//首先加载环境变量
	_ = LoadEnv()
}

func LoadEnv() error {
	env := os.Getenv(ENV)
	if env != "" {
		return json.Unmarshal([]byte(env), &options)
	}
	return nil
}

func StoreEnv() error {
	cfg, err := json.Marshal(&options)
	if err != nil {
		return err
	}
	return os.Setenv(ENV, string(cfg))
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Load(configure, &options)
}
