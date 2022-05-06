package master

import "github.com/zgwit/iot-master/connect"

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

	err = LoadTunnels()
	if err != nil {
		return err
	}

	err = LoadPipes()
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
	allPipes.Range(func(key, value interface{}) bool {
		dev := value.(*Pipe)
		_ = dev.Close()
		return true
	})
	allTunnels.Range(func(key, value interface{}) bool {
		tnl := value.(connect.Tunnel)
		_ = tnl.Close()
		return true
	})
}
