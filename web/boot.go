package web

import "github.com/zgwit/iot-master/v4/boot"

func init() {
	boot.Register("web", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config"},
	})
}
