package internal

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"sync"
)

var allDevices sync.Map
var allProjects sync.Map

func Start() error {
	err := LoadDevices()
	if err != nil {
		return err
	}
	err = LoadProjects()
	return err
}

func LoadDevices() error {
	devices := make([]*Device, 0)
	err := database.Device.All(devices)
	if err == storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}
	for _, d := range devices {
		allDevices.Store(d.Id, d)
		if d.Disabled {
			continue
		}

		err = d.Init()
		if err != nil {
			//TODO log
		} else {
			err = d.Start()
			if err != nil {
				//TODO log
			}
		}
	}
	return nil
}

func LoadDevice(id int) (*Device, error) {
	device := &Device{}
	err := database.Device.One("id", id, device)
	if err != nil {
		return nil, err
	}
	allDevices.Store(id, device)

	err = device.Init()
	if err != nil {
		return nil, err
	} else {
		err = device.Start()
		if err != nil {
			return nil, err
		}
	}
	return device, nil
}

func GetDevice(id int) *Device {
	d, ok := allDevices.Load(id)
	if ok {
		return d.(*Device)
	}
	return nil
}

func RemoveDevice(id int) error {
	d, ok := allDevices.LoadAndDelete(id)
	if ok {
		dev := d.(*Device)
		return dev.Stop()
	}
	return nil //error
}

func GetProject(id int) *Project {
	d, ok := allProjects.Load(id)
	if ok {
		return d.(*Project)
	}
	return nil
}

func RemoveProject(id int) error {
	d, ok := allProjects.LoadAndDelete(id)
	if ok {
		dev := d.(*Project)
		return dev.Stop()
	}
	return nil //error
}

func LoadProjects() error {
	project := make([]*Project, 0)
	err := database.Project.All(project)
	if err == storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}
	for _, p := range project {
		allProjects.Store(p.Id, project)
		if p.Disabled {
			continue
		}

		err = p.Init()
		if err != nil {
			//TODO log
		} else {
			err = p.Start()
			if err != nil {
				//TODO log
			}
		}
	}
	return nil
}

func LoadProject(id int) (*Project, error) {
	project := &Project{}
	err := database.Project.One("id", id, project)
	if err != nil {
		return nil, err
	}
	allProjects.Store(id, project)

	err = project.Init()
	if err != nil {
		return nil, err
	} else {
		err = project.Start()
		if err != nil {
			return nil, err
		}
	}
	return project, nil
}
