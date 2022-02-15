package master

import (
	"fmt"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/master/calc"
	"github.com/zgwit/iot-master/protocol"
)

type Adapter struct {
	slave    int
	protocol protocol.Protocol
	points   []Point

	events.EventEmitter
}

func (a *Adapter) Set(key string, value float64) error {
	for _, p := range a.points {
		if p.Name == key {
			data := p.Type.Encode(value, p.LittleEndian)
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
			v, err := p.Type.Decode(b, p.LittleEndian)
			if err != nil {
				return 0, err
			}
			//go func
			a.Emit("data", calc.Context{key: v})
			return v, nil
		}
	}

	return 0, fmt.Errorf("Unknown point %s ", key)
}

func (a *Adapter) Read(code, address, length int) (calc.Context, error) {
	//读取数据
	buf, err := a.protocol.Read(a.slave, code, address, length)
	if err != nil {
		return nil, err
	}

	//解析数据
	ctx := make(calc.Context)
	for _, p := range a.points {
		if address <= p.Address && p.Address < address+length {
			v, err := p.Type.Decode(buf[p.Address-p.Address:], p.LittleEndian)
			if err != nil {
				return nil, err
			}
			ctx[p.Name] = v
		}
	}
	//TODO 放这里不太合适
	a.Emit("data", ctx)

	return ctx, nil
}
