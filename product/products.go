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
	var m types.Product
	has, err := db.Engine.ID(id).Get(&m)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到设备%s", id)
	}

	return From(&m)
}

func From(product *types.Product) error {
	p := New(product)

	products.Store(product.Id, p)

	//err := db.Engine.Where("product_id = ?", id).And("disabled = ?", false).Find(&p.ExternalValidators)
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Engine.Where("product_id = ?", id).And("disabled = ?", false).Find(&p.ExternalAggregators)
	//if err != nil {
	//	return err
	//}

	return nil
}

func LoadAll() error {
	//开机加载所有产品，好像没有必要???
	var ps []*types.Product
	err := db.Engine.Find(&ps)
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
