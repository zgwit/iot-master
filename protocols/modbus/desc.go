package modbus

import "github.com/zgwit/iot-master/protocol"

//TODO const
var DescRTU = protocol.Item{
	Name:    "ModbusRTU",
	Version: "1.0",
	Label:   "Modbus RTU",
	Factory: newRTU,
}

var DescTCP = protocol.Item{
	Name:    "ModbusTCP",
	Version: "1.0",
	Label:   "Modbus TCP",
	Factory: newTCP,
}
