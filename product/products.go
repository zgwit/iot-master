package product

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/types"
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
	fn := fmt.Sprintf("product/%s/manifest.yaml", id)

	var m Manifest
	err := lib.LoadYaml(fn, &m)
	if err != nil {
		return err
	}

	return From(&m)
}

func From(product *Manifest) error {
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
	var ps []*types.Product
	err := db.Engine.Cols("id", "disabled").Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = Load(p.Id)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
