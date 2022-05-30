package siemens

import (
	"fmt"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocols/protocol"
	"regexp"
	"strconv"
)

type Address struct {
	Code   byte
	Block  uint16
	Offset uint32
}

func (a *Address) String() string {
	code := ""
	switch a.Code {
	case 0x81:
		code = "I"
	case 0x82:
		code = "Q"
	case 0x83:
		code = "M"
	case 0x84:
		code = "D" //DB V
	case 0x1D:
		code = "T"
	case 0x1C:
		code = "C"
	}
	return code + strconv.Itoa(int(a.Offset))
}

func (a *Address) Resolve(data []byte, from protocol.Addr, tp model.DataType, le bool, precision int) (interface{}, bool) {
	base := from.(*Address)
	if base.Code != a.Code {
		return nil, false
	}
	if base.Block != a.Block {
		return nil, false
	}
	cursor := int(a.Offset - base.Offset)
	if cursor < 0 || cursor > len(data) {
		return nil, false
	}
	val, err := tp.Decode(data[cursor:], le, precision)
	if err != nil {
		return nil, false
	}
	return val, true
}

var addrRegexp *regexp.Regexp

func init() {
	addrRegexp = regexp.MustCompile(`^(I|Q|M|D|DB|V|T|C)(\d+)(.\d+)?$`)
}

func ParseAddress(addr string) (protocol.Addr, error) {
	ss := addrRegexp.FindStringSubmatch(addr)
	if ss == nil || len(ss) < 3 {
		return nil, fmt.Errorf("西门子地址格式错误: %s", addr)
	}
	var code uint8 = 1
	switch ss[1] {
	case "I":
		code = 0x81
	case "Q":
		code = 0x82
	case "M":
		code = 0x83
	case "D":
		code = 0x84
	case "DB":
		code = 0x84
	case "T":
		code = 0x1D
	case "C":
		code = 0x1C
	case "V":
		code = 0x84
	}
	offset, _ := strconv.ParseUint(ss[2], 10, 16)
	//offset, _ := strconv.Atoi(ss[2])
	address := &Address{
		Code:   code,
		Offset: uint32(offset),
	}
	if len(ss) == 4 {
		address.Block = uint16(address.Offset)
		offset, _ = strconv.ParseUint(ss[3][1:], 10, 16)
		address.Offset = uint32(offset)
	}
	return address, nil
}
