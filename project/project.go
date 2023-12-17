package project

import (
	"github.com/zgwit/iot-master/v4/internal/device"
	"github.com/zgwit/iot-master/v4/types"
)

type Project struct {
	*Manifest

	Devices map[string]*device.Device

	ExternalValidators  []*types.ExternalValidator
	ExternalAggregators []*types.ExternalAggregator
}

func New(manifest *Manifest) *Project {
	return &Project{
		Manifest: manifest,
		Devices:  make(map[string]*device.Device),
		//Values: map[string]float64{},
	}
}
