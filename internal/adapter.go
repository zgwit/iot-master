package interval

import (
	"fmt"
	"github.com/zgwit/iot-master/protocol"
)

type Adapter struct {
	slave int
	protocol protocol.Protocol
	points   []Point
}

func (a *Adapter) Set(key string, value float64) error {
	for _, p := range a.points {
		if p.Name == key {
			data := p.Type.Encode(value)
			return a.protocol.Write(a.slave, p.Code, p.Address, data)
		}
	}

	return fmt.Errorf("Unknown point %s ", key)
}

func (a *Adapter) Get(key string) (float64, error) {
	for _, p := range a.points {
		if p.Name == key {
			//使用立即读
			b, err := a.protocol.ImmediateRead(a.slave, p.Code, p.Address, p.Type.Size())
			if err != nil {
				return 0, err
			}
			//解析数据
			v, err := p.Type.Decode(b)
			if err != nil {
				return 0, err
			}
			return v, nil
		}
	}

	return 0, fmt.Errorf("Unknown point %s ", key)
}

func (a *Adapter) Read(code, address, length int) (Context, error) {

	return nil, nil
}
