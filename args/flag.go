package args

import (
	"flag"
	"github.com/zgwit/iot-master/v4/build"
	"os"
)

var (
	showHelp    bool
	showVersion bool
	Install     bool
	Uninstall   bool
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&Install, "i", false, "Install service")
	flag.BoolVar(&Uninstall, "u", false, "Uninstall service")
}

func Parse() {
	flag.Parse()
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}
	if showVersion {
		build.Print()
		os.Exit(0)
	}
}
