package project

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
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
	fn := fmt.Sprintf("product/%s/manifest.yaml", id)

	var m Manifest
	err := lib.LoadYaml(fn, &m)
	if err != nil {
		return err
	}
	return From(&m)
}

func Store(id string, m *Manifest) error {
	fn := fmt.Sprintf("project/%s/manifest.yaml", id)
	err := lib.StoreYaml(fn, m)
	if err != nil {
		return err
	}
	return From(m)
}

func From(project *Manifest) error {
	p := New(project)

	projects.Store(project.Id, p)

	err := db.Engine.Where("project_id = ?", project.Id).And("disabled = ?", false).Find(&p.ExternalValidators)
	if err != nil {
		return err
	}

	err = db.Engine.Where("project_id = ?", project.Id).And("disabled = ?", false).Find(&p.ExternalAggregators)
	if err != nil {
		return err
	}

	return nil
}

func LoadAll() error {
	//开机加载所有产品，好像没有必要???
	var ps []*types.Project
	err := db.Engine.Cols("id", "disabled").Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		if p.Disabled {
			continue
		}
		err = Load(p.Id)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
