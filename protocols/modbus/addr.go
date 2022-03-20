package modbus

import (
	"errors"
	"github.com/zgwit/iot-master/protocol"
	"regexp"
	"strconv"
)

type Address struct {
	Slave  uint8  `json:"slave"`
	Code   uint8  `json:"code"`
	Offset uint16 `json:"offset"`
}

func (a *Address) String() string {

	return ""
}

var addrRegexp *regexp.Regexp
func init() {
	addrRegexp, _ = regexp.Compile(`^(X|D|O)(\d+)$`)

}

func ParseAddress(add string) (protocol.Address, error){
	ss := addrRegexp.FindStringSubmatch(add)
	if ss ==  nil || len(ss) != 3 {
		return nil, errors.New("unknown address")
	}
	var code uint8 = 1
	switch ss[1] {
	case "D": code = 1
	case "I": code = 2
	}
	offset, _ := strconv.ParseUint(ss[2], 10, 10)
	//offset, _ := strconv.Atoi(ss[2])
	return &Address{
		Code: code,
		Offset: uint16(offset),
	}, nil
}

// TODO const
var DescRTU = protocol.Item{
	Name:    "ModbusRTU",
	Version: "1.0",
	Label:   "Modbus RTU",
	Factory: newRTU,
	Address: ParseAddress,
}

var DescTCP = protocol.Item{
	Name:    "ModbusTCP",
	Version: "1.0",
	Label:   "Modbus TCP",
	Factory: newTCP,
	Address: ParseAddress,
}
