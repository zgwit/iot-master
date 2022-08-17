package config

import (
	"iot-master/helper"
)

//History 参数
type History struct {
	Type    string         `yaml:"type" json:"type"`
	Options helper.Options `yaml:"options" json:"options"`
}

var HistoryDefault = History{
	Type:    "embed",
	Options: helper.Options{"data_path": "history"},
}
