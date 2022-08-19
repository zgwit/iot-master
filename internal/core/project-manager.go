package core

import (
	"iot-master/internal/db"
	"iot-master/internal/log"
	"iot-master/model"
	"sync"
)

var allProjects sync.Map


func GetProject(id uint64) *Project {
	d, ok := allProjects.Load(id)
	if ok {
		return d.(*Project)
	}
	return nil
}

func RemoveProject(id uint64) error {
	d, ok := allProjects.LoadAndDelete(id)
	if ok {
		dev := d.(*Project)
		return dev.Stop()
	}
	return nil //error
}


func LoadProjects() error {
	return db.Store().ForEach(nil, func(p *model.Project) error {
		if p.Disabled {
			return nil
		}

		prj, err := NewProject(p)
		if err != nil {
			log.Error(err)
			return nil
		}
		allProjects.Store(p.Id, prj)

		err = prj.Start()
		if err != nil {
			log.Error(err)
		}
		return nil
	})
}


func LoadProject(id uint64) (*Project, error) {
	var project model.Project
	err := db.Store().Get(id, &project)
	if err != nil {
		return nil, err
	}

	prj, err := NewProject(&project)
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
