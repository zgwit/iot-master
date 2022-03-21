package master

import (
	"fmt"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
)

//Adapter 数据解析器（可能要改名）
type Adapter struct {
	slave    int
	protocol protocol.Protocol
	points   []Point

	events.EventEmitter
}

//Set 写数据位
func (a *Adapter) Set(key string, value float64) error {
	for _, p := range a.points {
		if p.Name == key {
			data := p.Type.Encode(value, p.LittleEndian)
			return a.protocol.Write(p.Addr, data)
		}
	}

	return fmt.Errorf("Unknown point %s ", key)
}

//Get 读数据位
func (a *Adapter) Get(key string) (float64, error) {

	for _, p := range a.points {
		if p.Name == key {
			//使用立即读
			b, err := a.protocol.Immediate(p.Addr, uint16(p.Type.Size()))
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

//Read 读多数据
func (a *Adapter) Read(addr protocol.Addr, length uint16) (calc.Context, error) {
	//读取数据
	buf, err := a.protocol.Read(addr, length)
	if err != nil {
		return nil, err
	}

	//解析数据
	ctx := make(calc.Context)
	for _, p := range a.points {
		if addr <= p.Address && p.Address < addr+length {
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
