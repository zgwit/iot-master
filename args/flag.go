package args

import (
	"flag"
	"os"
	"path/filepath"
)

var (
	help       bool
	ConfigPath string
	Install    bool
	Uninstall  bool
)

func init() {
	app, _ := filepath.Abs(os.Args[0])
	//log.Println("app.path", app)
	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&ConfigPath, "c", app+".yaml", "Configure path")
	flag.BoolVar(&Install, "i", false, "Install service")
	flag.BoolVar(&Uninstall, "u", false, "Uninstall service")
}

func Parse() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
}
