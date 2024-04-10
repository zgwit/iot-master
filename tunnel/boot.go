package tunnel

import (
	"github.com/zgwit/iot-master/v4/boot"
	"github.com/zgwit/iot-master/v4/log"
)

func init() {
	boot.Register("tunnel", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"database"},
	})
}

func Startup() error {
	err := LoadSerials()
	if err != nil {
		return err
	}
	err = LoadClients()
	if err != nil {
		return err
	}
	err = LoadServers()
	if err != nil {
		return err
	}
	return nil
}

func Shutdown() error {
	links.Range(func(name string, link *Link) bool {
		err := link.Close()
		if err != nil {
			log.Error(err)
		}
		return true
	})
	return nil
}
