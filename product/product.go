package product

import "github.com/zgwit/iot-master/v4/types"

type Product struct {
	*Model

	ExternalValidators  []*types.Validator
	ExternalAggregators []*types.Aggregator
}

func New(model *Model) *Product {
	return &Product{
		Model: model,
		//Values: map[string]float64{},
	}
}
