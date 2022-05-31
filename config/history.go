package config

import (
	"github.com/zgwit/iot-master/helper"
)

//History 参数
type History struct {
	Type    string         `yaml:"type"`
	Options helper.Options `yaml:"options"`
}

var HistoryDefault = History{
	Type:    "embed",
	Options: helper.Options{"data_path": "history"},
}
