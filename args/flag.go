package args

import (
	"flag"
	"os"
)

var (
	showHelp  bool
	Install   bool
	Uninstall bool
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "帮助")
	flag.BoolVar(&Install, "i", false, "安装服务")
	flag.BoolVar(&Uninstall, "u", false, "卸载服务")
}

func Parse() {
	flag.Parse()
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}
}
