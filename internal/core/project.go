package core

import "github.com/zgwit/iot-master/v2/model"

func NewProject(id string) *Project {
	return &Project{
		Id:      id,
		Values:  make(map[string]any),
		Devices: make(map[string]*Device),
	}
}

type Project struct {
	Id      string
	Values  map[string]any
	Status  model.Status
	Devices map[string]*Device
}

func (p *Project) Assign(points map[string]any) error {

	return nil
}

func (p *Project) Refresh() error {

	return nil
}

//func (p *Project) Status() error {
//
//	return nil
//}
