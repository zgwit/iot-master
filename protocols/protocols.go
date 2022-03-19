package protocols

import (
	"fmt"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/protocols/modbus"
)

var protocols = []protocol.Desc{
	modbus.DescRTU,
	modbus.DescTCP,
}

func Protocols() []protocol.Desc {
	return protocols
}

func Create(name string, options interface{}) (protocol.Protocol, error) {
	for _, d := range protocols {
		if d.Name == name {
			return d.Factory(options), nil
		}
	}
	return nil, fmt.Errorf("unkown protocol: %s", name)
}
