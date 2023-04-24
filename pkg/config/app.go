package config

import (
	"os"
	"path/filepath"
	"strings"
)

func AppName() string {
	//插件的args参数为空
	if len(os.Args) == 0 {
		return "iot-master"
	}
	app, _ := filepath.Abs(os.Args[0])
	ext := filepath.Ext(os.Args[0])
	return strings.TrimSuffix(app, ext)
}
