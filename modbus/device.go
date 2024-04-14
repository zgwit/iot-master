package modbus

import (
	"github.com/zgwit/iot-master/v4/db"
)

func init() {
	db.Register(new(Device))
}

type Station struct {
	Slave uint8 `json:"slave"`
}

type Device struct {
	Id string `json:"id" xorm:"pk"`

	//modbus站号
	//ModbusStation uint8 `json:"modbus_station,omitempty"`
	Station Station `json:"station,omitempty" xorm:"json"`

	//映射和轮询表
	pollers *[]*Poller
	mappers *[]*Mapper
}

func (p *Device) Lookup(name string) *Mapper {
	for _, m := range *p.mappers {
		if m.Name == name {
			return m
		}
	}
	return nil
}
