package modbus

import (
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/lib"
	"os"
	"path/filepath"
)

type Filter struct {
	Name       string `json:"name"`       //字段
	Expression string `json:"expression"` //表达式
	//Entire     bool   `json:"entire"`
}

type Calculator struct {
	Name       string `json:"name"`       //赋值
	Expression string `json:"expression"` //表达式
}

type Product struct {
	Pollers []*Poller `json:"pollers"`
	Mappers []*Mapper `json:"mappers"`
	//Filters     []*Filter     `json:"filters"`
	//Calculators []*Calculator `json:"calculators"`
	//Created     time.Time     `json:"created"` //创建时间
}

func (p *Product) Lookup(name string) *Mapper {
	for _, m := range p.Mappers {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (p *Poller) Parse(mappers []*Mapper, buf []byte, values map[string]any) error {
	for _, m := range mappers {
		if p.Code == m.Code &&
			p.Address <= m.Address &&
			p.Length > m.Address-p.Address {
			ret, err := m.Parse(p.Address, buf)
			if err != nil {
				//log.Error(err)
				return err
			}
			//03指令 的 位类型
			if rets, ok := ret.(map[string]bool); ok {
				for k, v := range rets {
					values[k] = v
				}
			} else {
				values[m.Name] = ret
			}
		}
	}
	return nil
}

var products lib.Map[Product]

func GetProduct(id, version string) (*Product, error) {
	p := products.Load(id + version)
	if p == nil {
		err := LoadProduct(id, version)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

func GetProduct2(id, version string) (*Product, error) {
	p := products.Load(id + version)
	if p == nil {
		return nil, errors.New("找不到产品")
	}
	return p, nil
}

func LoadProduct(id, version string) error {
	//TODO 处理 loading状态

	var product Product
	products.Store(id+version, &product)

	fn := filepath.Join(viper.GetString("data"), "product", id, version, "pollers.json")
	buf, err := os.ReadFile(fn)
	if err == nil {
		err = json.Unmarshal(buf, &product.Pollers)
		if err != nil {
			return err
		}
	}

	fn = filepath.Join(viper.GetString("data"), "product", id, version, "mappers.json")
	buf, err = os.ReadFile(fn)
	if err == nil {
		err = json.Unmarshal(buf, &product.Mappers)
		if err != nil {
			return err
		}
	}

	return nil
}
