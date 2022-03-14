package args

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

var (
	help       bool
	ConfigPath string
	Install    bool
	Uninstall  bool
)

func init() {
	app, _ := filepath.Abs(os.Args[0])
	ext := filepath.Ext(os.Args[0])
	//替换后缀名.exe为.yaml
	cfg := strings.TrimSuffix(app, ext) + ".yaml"

	//log.Println("app.path", app)
	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&ConfigPath, "c", cfg, "Configure path")
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
