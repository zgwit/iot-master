package product

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
)

var products lib.Map[Product]

func Ensure(id string) (*Product, error) {
	dev := products.Load(id)
	if dev == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		dev = products.Load(id)
	}
	return dev, nil
}

func Get(id string) *Product {
	return products.Load(id)
}

func Load(id string) error {
	var p Model
	get, err := db.Engine.ID(id).Get(&p)
	if err != nil {
		return err
	}
	if !get {
		return fmt.Errorf("product %s not found", id)
	}

	return From(&p)
}

func From(product *Model) error {
	p := New(product)

	products.Store(product.Id, p)

	err := db.Engine.Where("product_id = ?", product.Id).And("disabled = ?", false).Find(&p.ExternalValidators)
	if err != nil {
		return err
	}

	err = db.Engine.Where("product_id = ?", product.Id).And("disabled = ?", false).Find(&p.ExternalAggregators)
	if err != nil {
		return err
	}

	return nil
}

func LoadAll() error {
	//开机加载所有产品，好像没有必要???
	var ps []*Model
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = From(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
