package internal

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

var Products lib.Map[Product]

type Product struct {
	model  *model.Product
	eval   []gval.Evaluable
	values map[string]float64
}

func LoadProduct(product *model.Product) error {
	//log.Info("load product", product.Id, product.Name)
	p := &Product{
		model:  product,
		values: map[string]float64{},
	}
	for _, c := range product.Constraints {
		eval, err := calc.New(c.Expression)
		if err != nil {
			return err
		}
		p.eval = append(p.eval, eval)
	}
	for _, param := range product.Parameters {
		p.values[param.Name] = param.Default
	}

	Products.Store(product.Id, p)

	return nil
}

func LoadProducts() error {
	//开机加载所有产品，好像没有必要???
	var ps []*model.Product
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = LoadProduct(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
