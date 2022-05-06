package master

import (
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"sync"
)

var allDevices sync.Map
var allProjects sync.Map


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
			log.Error(err)
			continue
		}
		allDevices.Store(d.Id, dev)

		err = dev.Start()
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

//LoadDevice 加载设备
func LoadDevice(id int) (*Device, error) {
	device := &model.Device{}
	err := database.Master.One("Id", id, device)
	if err != nil {
		return nil, err
	}

	dev, err := NewDevice(device)
	if err != nil {
		return dev, err
	}
	//allDevices.Store(device.Id, dev)
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
		if p.Disabled {
			continue
		}

		prj, err := NewProject(p)
		if err != nil {
			log.Error(err)
			continue
		}
		allProjects.Store(p.Id, prj)

		err = prj.Start()
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

//LoadProject 加载项目
func LoadProject(id int) (*Project, error) {
	project := &model.Project{}
	err := database.Master.One("Id", id, project)
	if err != nil {
		return nil, err
	}

	prj, err := NewProject(project)
	if err != nil {
		return nil, err
	}

	allProjects.Store(id, prj)

	err = prj.Start()
	if err != nil {
		return nil, err
	}
	return prj, nil
}
