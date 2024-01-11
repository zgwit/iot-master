package project

import (
	"github.com/zgwit/iot-master/v4/internal/device"
	"github.com/zgwit/iot-master/v4/types"
)

type Project struct {
	*types.Project

	Devices map[string]*device.Device

	ExternalValidators  []*types.ExternalValidator
	ExternalAggregators []*types.ExternalAggregator
}

func New(project *types.Project) *Project {
	return &Project{
		Project: project,
		Devices: make(map[string]*device.Device),
		//Values: map[string]float64{},
	}
}
