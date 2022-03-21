package master

import (
	"fmt"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
)

//Mapper 数据解析器（可能要改名）
type Mapper struct {
	events.EventEmitter

	slave int //TODO 从站号

	protocol protocol.Protocol
	points   []Point
}

func newMapper(points []model.Point, protocol protocol.Protocol) *Mapper {
	adapter := &Mapper{
		protocol: protocol,
		points:   make([]Point, len(points)),
	}
	for i, p := range points {
		addr, _ := protocol.Address(p.Address)
		adapter.points[i] = Point{
			Point: p,
			Addr:  addr,
		}
	}
	return adapter
}

//Set 写数据位
func (m *Mapper) Set(key string, value float64) error {
	for _, p := range m.points {
		if p.Name == key {
			data := p.Type.Encode(value, p.LittleEndian)
			return m.protocol.Write(p.Addr, data)
		}
	}

	return fmt.Errorf("Unknown point %s ", key)
}

//Get 读数据位
func (m *Mapper) Get(key string) (float64, error) {

	for _, p := range m.points {
		if p.Name == key {
			//使用立即读
			b, err := m.protocol.Immediate(p.Addr, uint16(p.Type.Size()))
			if err != nil {
				return 0, err
			}
			//解析数据
			v, err := p.Type.Decode(b, p.LittleEndian)
			if err != nil {
				return 0, err
			}
			//go func
			m.Emit("data", calc.Context{key: v})
			return v, nil
		}
	}

	return 0, fmt.Errorf("Unknown point %s ", key)
}

//Read 读多数据
func (m *Mapper) Read(addr protocol.Addr, length int) (calc.Context, error) {
	//读取数据
	buf, err := m.protocol.Read(addr, uint16(length))
	if err != nil {
		return nil, err
	}

	//解析数据
	ctx := make(calc.Context)
	for _, p := range m.points {
		offset := p.Addr.Diff(addr)
		if offset > 0 && offset < length {
			v, err := p.Type.Decode(buf[offset:], p.LittleEndian)
			if err != nil {
				return nil, err
			}
			ctx[p.Name] = v
		}
	}
	//TODO 放这里不太合适
	m.Emit("data", ctx)

	return ctx, nil
}
