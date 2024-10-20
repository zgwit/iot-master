package web

import "github.com/zgwit/iot-master/boot"

func init() {
	boot.Register("web", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config"},
	})
}
