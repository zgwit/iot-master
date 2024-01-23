package space

import (
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/types"
)

type Space struct {
	id   string
	name string

	values map[string]any

	devices map[string]*device.Device
}

func (s *Space) PutDevice(name string, dev *device.Device) {
	s.devices[name] = dev
	s.values[name] = dev.Values()

	dev.EventData.On(func(value map[string]any) {
		//此处用来触发情景模式
	})

}

func New(space *types.Space) *Space {
	return &Space{
		id:      space.Id,
		name:    space.Name,
		devices: make(map[string]*device.Device),
		values:  make(map[string]any),
	}
}
