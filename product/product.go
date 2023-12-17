package product

import (
	"github.com/zgwit/iot-master/v4/types"
)

type Product struct {
	*Manifest

	ExternalValidators  []*types.ExternalValidator
	ExternalAggregators []*types.ExternalAggregator
}

func New(manifest *Manifest) *Product {
	return &Product{
		Manifest: manifest,
		//Values: map[string]float64{},
	}
}
