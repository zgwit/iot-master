package modbus

import "github.com/zgwit/iot-master/protocol"

// Function Code
const (
	// Bit access
	FuncCodeReadDiscreteInputs = 2
	FuncCodeReadCoils          = 1
	FuncCodeWriteSingleCoil    = 5
	FuncCodeWriteMultipleCoils = 15

	// 16-bit access
	FuncCodeReadInputRegisters         = 4
	FuncCodeReadHoldingRegisters       = 3
	FuncCodeWriteSingleRegister        = 6
	FuncCodeWriteMultipleRegisters     = 16
	FuncCodeReadWriteMultipleRegisters = 23
	FuncCodeMaskWriteRegister          = 22
	FuncCodeReadFIFOQueue              = 24
	FuncCodeOtherReportSlaveID         = 17
	// FuncCodeDiagReadException          = 7
	// FuncCodeDiagDiagnostic             = 8
	// FuncCodeDiagGetComEventCnt         = 11
	// FuncCodeDiagGetComEventLog         = 12

)

// Exception Code
const (
	ExceptionCodeIllegalFunction                    = 1
	ExceptionCodeIllegalDataAddress                 = 2
	ExceptionCodeIllegalDataValue                   = 3
	ExceptionCodeServerDeviceFailure                = 4
	ExceptionCodeAcknowledge                        = 5
	ExceptionCodeServerDeviceBusy                   = 6
	ExceptionCodeNegativeAcknowledge                = 7
	ExceptionCodeMemoryParityError                  = 8
	ExceptionCodeGatewayPathUnavailable             = 10
	ExceptionCodeGatewayTargetDeviceFailedToRespond = 11
)

var Codes = []protocol.Code{
	{"C", "线圈"},
	{"D", "离散输入"},
	{"H", "保持寄存器"},
	{"I", "输入寄存器"},
}

var DescRTU = protocol.Protocol{
	Name:    "ModbusRTU",
	Version: "1.0",
	Label:   "Modbus RTU",
	Codes:   Codes,
	Factory: NewRTU,
}

var DescTCP = protocol.Protocol{
	Name:    "ModbusTCP",
	Version: "1.0",
	Label:   "Modbus TCP",
	Codes:   Codes,
	Factory: NewTCP,
}
