package broker

import "github.com/zgwit/iot-master/v4/boot"

func init() {
	boot.Register("broker", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"database"},
	})
}
