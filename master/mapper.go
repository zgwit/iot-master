package master

import (
	"fmt"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
	"math"
)

//Mapper 数据解析器（可能要改名）
type Mapper struct {
	//model.Mapping
	events.EventEmitter

	station int //从站号

	adapter protocol.Adapter
	points  []Point
}

func newMapper(station int, points []*model.Point, adapter protocol.Adapter) (*Mapper, error) {
	mapper := &Mapper{
		station: station,
		adapter: adapter,
		points:  make([]Point, len(points)),
	}
	for i, p := range points {
		addr, err := adapter.Address(p.Address)
		if err != nil {
			return nil, err
		}
		mapper.points[i] = Point{
			Point: *p,
			Addr:  addr,
		}
	}
	return mapper, nil
}

//Set 写数据位
func (m *Mapper) Set(key string, value float64) error {
	for _, p := range m.points {
		if p.Name == key {
			data := p.Type.Encode(value, p.LittleEndian)
			return m.adapter.Write(m.station, p.Addr, data)
		}
	}

	return fmt.Errorf("Unknown point %s ", key)
}

//Get 读数据位
func (m *Mapper) Get(key string) (float64, error) {

	for _, p := range m.points {
		if p.Name == key {
			//使用立即读
			b, err := m.adapter.Immediate(m.station, p.Addr, p.Type.Size())
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
	buf, err := m.adapter.Read(m.station, addr, length)
	if err != nil {
		return nil, err
	}

	//解析数据
	ctx := make(calc.Context)
	for _, p := range m.points {
		offset := p.Addr.Diff(addr)
		if offset >= 0 && offset < length {
			v, err := p.Type.Decode(buf[offset*2:], p.LittleEndian)
			if err != nil {
				return nil, err
			}
			//倍率计算
			if p.Precision > 0 {
				v /= math.Pow10(p.Precision)
			}
			ctx[p.Name] = v
		}
	}
	//TODO 放这里不太合适
	m.Emit("data", ctx)

	return ctx, nil
}
