package args

import (
	"flag"
	"github.com/zgwit/iot-master/v4/pkg/build"
	"os"
	"path/filepath"
	"strings"
)

var (
	showHelp    bool
	showVersion bool
	ConfigPath  string
	Install     bool
	Uninstall   bool
)

func init() {
	//参数配置
	app, _ := filepath.Abs(os.Args[0])
	ext := filepath.Ext(os.Args[0])
	//替换后缀名.exe为.yaml
	cfg := strings.TrimSuffix(app, ext) + ".yaml"

	//log.Println("app.path", app)
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.StringVar(&ConfigPath, "c", cfg, "Configure path")
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
