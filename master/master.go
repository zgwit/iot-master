package master

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"sync"
)

var allTunnels sync.Map
var allLinks sync.Map
var allDevices sync.Map
var allProjects sync.Map

//Start 启动
func Start() error {
	err := LoadTunnels()
	if err != nil {
		return err
	}

	err = LoadDevices()
	if err != nil {
		return err
	}

	err = LoadProjects()
	return err
}

func Stop() {
	allTunnels.Range(func(key, value interface{}) bool {
		tnl := value.(connect.Tunnel)
		_ = tnl.Close()
		return true
	})
	allDevices.Range(func(key, value interface{}) bool {
		dev := value.(*Device)
		_ = dev.Stop()
		return true
	})
	allProjects.Range(func(key, value interface{}) bool {
		prj := value.(*Project)
		_ = prj.Stop()
		return true
	})
}


//LoadTunnels 加载通道
func LoadTunnels() error {
	var tunnels []*model.Tunnel
	err := database.Master.All(&tunnels)
	if err == storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}
	for _, d := range tunnels {
		if d.Disabled {
			continue
		}

		tnl, err := connect.NewTunnel(d)
		if err != nil {
			//return err
			//TODO log
			continue
		}
		allTunnels.Store(d.ID, tnl)

		err = tnl.Open()
		if err != nil {
			//TODO log
		}

		tnl.On("link", func(link connect.Link) {
			allLinks.Store(link.ID(), link)
			//TODO 找到相关Device
		})
	}
	return nil
}


//GetTunnel 获取通道
func GetTunnel(id int) connect.Tunnel {
	d, ok := allTunnels.Load(id)
	if ok {
		return d.(connect.Tunnel)
	}
	return nil
}



//LoadDevices 加载设备
func LoadDevices() error {
	var devices []*model.Device
	err := database.Master.All(&devices)
	if err == storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}
	for _, d := range devices {
		if d.Disabled {
			continue
		}

		dev, err := NewDevice(d)
		if err != nil {
			return err
		}
		allDevices.Store(d.ID, dev)

		err = dev.Start()
		if err != nil {
			//TODO log
		}
	}
	return nil
}

//LoadDevice 加载设备
func LoadDevice(id int) (*Device, error) {
	device := &model.Device{}
	err := database.Master.One("id", id, device)
	if err != nil {
		return nil, err
	}

	dev, err := NewDevice(device)
	if err != nil {
		return dev, err
	}
	//allDevices.Store(device.ID, dev)
	allDevices.Store(id, dev)
	err = dev.Start()
	return dev, nil
}

//GetDevice 获取设备
func GetDevice(id int) *Device {
	d, ok := allDevices.Load(id)
	if ok {
		return d.(*Device)
	}
	return nil
}

//RemoveDevice 删除设备
func RemoveDevice(id int) error {
	d, ok := allDevices.LoadAndDelete(id)
	if ok {
		dev := d.(*Device)
		return dev.Stop()
	}
	return nil //error
}

//GetProject 获取项目
func GetProject(id int) *Project {
	d, ok := allProjects.Load(id)
	if ok {
		return d.(*Project)
	}
	return nil
}

//RemoveProject 删除项目
func RemoveProject(id int) error {
	d, ok := allProjects.LoadAndDelete(id)
	if ok {
		dev := d.(*Project)
		return dev.Stop()
	}
	return nil //error
}

//LoadProjects 加载项目
func LoadProjects() error {
	var projects []*model.Project
	err := database.Master.All(&projects)
	if err == storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}
	for _, p := range projects {
		allProjects.Store(p.ID, projects)
		if p.Disabled {
			continue
		}

		prj, err := NewProject(p)
		if err != nil {
			//TODO log
			continue
		}
		err = prj.Init()
		if err != nil {
			//TODO log
			continue
		}
		err = prj.Start()
		if err != nil {
			//TODO log
		}
	}
	return nil
}

//LoadProject 加载项目
func LoadProject(id int) (*Project, error) {
	project := &model.Project{}
	err := database.Master.One("id", id, project)
	if err != nil {
		return nil, err
	}

	prj, err := NewProject(project)
	if err != nil {
		return nil, err
	}

	allProjects.Store(id, prj)

	err = prj.Init()
	if err != nil {
		return nil, err
	}
	err = prj.Start()
	if err != nil {
		return nil, err
	}
	return prj, nil
}
