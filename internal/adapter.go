package interval

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/zgwit/iot-master/protocol"
)

type Adapter struct {
	slave    int
	protocol protocol.Protocol
	points   []Point
	events   EventBus.Bus
}

func (a *Adapter) Init() {
	a.events = EventBus.New()
}

func (a *Adapter) Set(key string, value float64) error {
	//for _, p := range a.points {
	for i := 0; i < len(a.points); i++ {
		p := &a.points[i]
		if p.Name == key {
			data := p.Type.Encode(value, p.LittleEndian)
			return a.protocol.Write(a.slave, p.Code, p.Address, data)
		}
	}

	return fmt.Errorf("Unknown point %s ", key)
}

func (a *Adapter) Get(key string) (float64, error) {

	//for _, p := range a.points {
	for i := 0; i < len(a.points); i++ {
		p := &a.points[i]
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
			a.events.Publish("data", Context{key: v})
			return v, nil
		}
	}

	return 0, fmt.Errorf("Unknown point %s ", key)
}

func (a *Adapter) Read(code, address, length int) (Context, error) {
	//读取数据
	buf, err := a.protocol.Read(a.slave, code, address, length)
	if err != nil {
		return nil, err
	}

	//解析数据
	ctx := make(Context)
	for i := 0; i < len(a.points); i++ {
		p := &a.points[i]
		if address <= p.Address && p.Address < address+length {
			v, err := p.Type.Decode(buf[p.Address-p.Address:], p.LittleEndian)
			if err != nil {
				return nil, err
			}
			ctx[p.Name] = v
		}
	}
	//TODO 放这里不太合适
	a.events.Publish("data", ctx)

	return ctx, nil
}
