package space

import (
	"github.com/zgwit/iot-master/v4/boot"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
)

func init() {
	boot.Register("space", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"database"},
	})
}

func Startup() error {
	return nil
}

func Startup2() error {
	//开机加载所有空间，好像没有必要???

	var ps []*Space
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = From(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}

func Shutdown() error {
	//TODO 关闭空间

	return nil
}
