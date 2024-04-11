package user

import (
	"github.com/zgwit/iot-master/v4/boot"
	"github.com/zgwit/iot-master/v4/web"
)

func init() {
	boot.Register("user", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"database", "web"},
	})
}

func Startup() error {

	//鉴权接口
	web.Engine.GET("api/auth", auth)

	web.Engine.POST("api/login", login)

	return nil
}
