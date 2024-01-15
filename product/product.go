package product

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/types"
)

type Product struct {
	*types.Product
	*Manifest

	ExternalValidators  []*types.ExternalValidator
	ExternalAggregators []*types.ExternalAggregator
}

func (p *Product) StoreManifest() error {
	fn := fmt.Sprintf("%s/product/%s/manifest.yaml", viper.GetString("data"), p.Id)
	return lib.StoreYaml(fn, p.Manifest)
}

func New(product *types.Product) *Product {
	return &Product{
		Product: product,
		//values: map[string]float64{},
	}
}
