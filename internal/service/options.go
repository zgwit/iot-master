package service

import (
	"github.com/zgwit/iot-master/v4/config"
)

// Options 参数
type Options struct {
	Name         string            `json:"name,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Description  string            `json:"description,omitempty"`
	Executable   string            `json:"executable,omitempty"`
	Directory    string            `json:"directory,omitempty"`
	Arguments    []string          `json:"arguments,omitempty"`
	Dependencies []string          `json:"dependencies,omitempty"`
	Environments map[string]string `json:"environments,omitempty"`
}

func Default() Options {
	return Options{
		Name:        "iot-master",
		DisplayName: "物联大师",
		Description: "物联网数据中台",
	}
}

var options Options = Default()
var configure = "service"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
