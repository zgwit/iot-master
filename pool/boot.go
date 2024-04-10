package pool

import "github.com/zgwit/iot-master/v4/boot"

func init() {
	boot.Register("pool", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config"},
	})
}
