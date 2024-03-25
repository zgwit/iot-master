package project

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/space"
)

var projects lib.Map[Project]

func Get(id string) *Project {
	return projects.Load(id)
}

func Load(id string) error {
	var m Project
	has, err := db.Engine.ID(id).Get(&m)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到项目%s", id)
	}

	return From(&m)
}

func From(p *Project) error {
	projects.Store(p.Id, p)

	var ds []*space.Space
	err := db.Engine.Where("project_id=?", p.Id).Find(&ds)
	if err != nil {
		return err
	}

	for _, s := range ds {
		err := space.From(s)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func Boot() error {
	//开机加载所有项目，好像没有必要???

	var ps []*Project
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
