package project

import (
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/types"
)

type Manifest struct {
	types.ManifestBase `yaml:"inline"`

	//参数
	Parameters []*product.Parameter `json:"parameters,omitempty"` //参数

	Validators  []*types.Validator  `json:"validators,omitempty"`
	Aggregators []*types.Aggregator `json:"aggregators,omitempty"`
}
