package broker

import (
	"github.com/zgwit/iot-master/v4/boot"
	"github.com/zgwit/iot-master/v4/mqtt"
)

func init() {
	boot.Register("broker", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"database"},
	})

	//优先启动
	task := boot.Load("mqtt")
	task.Depends = append(task.Depends, "broker")

	//注册接口
	mqtt.CustomOpenConnectionFn = CustomOpenConnectionFn
}
