package internal

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

var Products lib.Map[Product]

type Product struct {
	model  *model.Product
	eval   []gval.Evaluable
	values map[string]float64
}

func Load(model *model.Product) error {
	log.Info("load product", model.Id, model.Name)
	pro := &Product{
		model:  model,
		values: map[string]float64{},
	}
	for _, c := range model.Constraints {
		eval, err := calc.New(c.Expression)
		if err != nil {
			return err
		}
		pro.eval = append(pro.eval, eval)
	}
	for _, p := range model.Parameters {
		pro.values[p.Name] = p.Default
	}

	Products.Store(model.Id, pro)

	return nil
}

func LoadProducts() error {
	//开机加载所有产品，好像没有必要

	return nil
}
