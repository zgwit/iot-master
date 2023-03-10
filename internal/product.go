package internal

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/lib"
)

var Products lib.Map[Product]

type Product struct {
	model *model.Product
	eval  []gval.Evaluable
}

func NewProduct(model *model.Product) *Product {
	pro := &Product{
		model: model,
		eval:  make([]gval.Evaluable, 0),
	}
	Products.Store(model.Id, pro)
	for _, c := range model.Constraints {
		eval, _ := calc.New(c.Expression)
		//TODO error
		pro.eval = append(pro.eval, eval)
	}
	return pro
}
