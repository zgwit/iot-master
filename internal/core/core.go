package core

import (
	"github.com/zgwit/iot-master/link"
)

//Start 启动
func Start() error {
	var err error
	err = LoadDevices()
	if err != nil {
		return err
	}

	err = LoadProjects()
	if err != nil {
		return err
	}

	err = LoadServers()
	if err != nil {
		return err
	}

	err = LoadTunnels()
	if err != nil {
		return err
	}

	return nil
}

func Stop() {
	allProjects.Range(func(key, value interface{}) bool {
		prj := value.(*Project)
		_ = prj.Stop()
		return true
	})
	allDevices.Range(func(key, value interface{}) bool {
		dev := value.(*Device)
		_ = dev.Stop()
		return true
	})
	allTunnels.Range(func(key, value interface{}) bool {
		tnl := value.(link.Tunnel)
		_ = tnl.Close()
		return true
	})
}
