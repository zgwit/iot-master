package config

import (
	"os"
	"path/filepath"
	"strings"
)

func AppName() string {
	app, _ := filepath.Abs(os.Args[0])
	ext := filepath.Ext(os.Args[0])
	return strings.TrimSuffix(app, ext)
}
