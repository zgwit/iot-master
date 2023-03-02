package args

import (
	"flag"
	"fmt"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	showHelp    bool
	showVersion bool
	ConfigPath  string
	Install     bool
	Uninstall   bool
)

var (
	Version   string
	GitHash   string
	BuildTime string
	Time      time.Time
)

func init() {

	//初始化编译时间
	if BuildTime != "" {
		var err error
		Time, err = http.ParseTime(BuildTime)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		Time = time.Now()
	}

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
		fmt.Printf("Version: %s \n", Version)
		fmt.Printf("Git Hash: %s \n", GitHash)
		fmt.Printf("Build Time: %s \n", BuildTime)
		os.Exit(0)
	}
}
