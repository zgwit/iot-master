package project

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/types"
)

var projects lib.Map[Project]

func Ensure(id string) (*Project, error) {
	dev := projects.Load(id)
	if dev == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		dev = projects.Load(id)
	}
	return dev, nil
}

func Get(id string) *Project {
	return projects.Load(id)
}

func Load(id string) error {
	var m types.Project
	has, err := db.Engine.ID(id).Get(&m)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到项目%s", id)
	}

	return From(&m)
}

func From(project *types.Project) error {
	p := New(project)

	projects.Store(project.Id, p)
	//
	//err := db.Engine.Where("project_id = ?", id).And("disabled = ?", false).Find(&p.ExternalValidators)
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Engine.Where("project_id = ?", id).And("disabled = ?", false).Find(&p.ExternalAggregators)
	//if err != nil {
	//	return err
	//}

	return nil
}

func Boot() error {
	//开机加载所有产品，好像没有必要???
	var ps []*types.Project
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = Load(p.Id)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
