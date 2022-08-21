package core

import (
	"github.com/timshannon/bolthold"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
	"sync"
)

var products sync.Map

func LoadProduct(id string) (*model.Product, error) {
	v, ok := products.Load(id)
	if ok {
		return v.(*model.Product), nil
	}

	//加载产品
	var product model.Product
	err := db.Store().Get(id, &product)
	if err == bolthold.ErrNotFound {
		return nil, err
	}

	products.Store(id, &product)

	return &product, nil
}
