package modbus

import "github.com/zgwit/iot-master/protocol"

type Address struct {
	Slave  uint8  `json:"slave"`
	Code   uint8  `json:"code"`
	Offset uint16 `json:"offset"`
}

func (a *Address) String() string {

	return ""
}

func AddressParser(add string) protocol.Address {

	return &Address{}
}

// TODO const
var DescRTU = protocol.Item{
	Name:    "ModbusRTU",
	Version: "1.0",
	Label:   "Modbus RTU",
	Factory: newRTU,
	AddressParser: AddressParser,
}

var DescTCP = protocol.Item{
	Name:    "ModbusTCP",
	Version: "1.0",
	Label:   "Modbus TCP",
	Factory: newTCP,
	AddressParser: AddressParser,
}
