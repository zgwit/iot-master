package modbus

import (
	"errors"
	"github.com/zgwit/iot-master/protocol"
	"regexp"
	"strconv"
)

type Address struct {
	Code   uint8  `json:"code"`
	Offset uint16 `json:"offset"`
}

func (a *Address) String() string {
	code := ""
	switch a.Code {
	case 1:
		code = "C"
	case 2:
		code = "D"
	case 3:
		code = "H"
	case 4:
		code = "I"
	}
	return code + strconv.Itoa(int(a.Offset))
}

func (a *Address) Diff(base protocol.Addr) int {
	start := base.(*Address)
	if start.Code != a.Code {
		return -1
	}
	return int(a.Offset - start.Offset)
}

var addrRegexp *regexp.Regexp

func init() {
	addrRegexp = regexp.MustCompile(`^(C|D|DI|H|I)(\d+)$`)
}

func ParseAddress(addr string) (protocol.Addr, error) {
	ss := addrRegexp.FindStringSubmatch(addr)
	if ss == nil || len(ss) != 3 {
		return nil, errors.New("unknown address")
	}
	var code uint8 = 1
	switch ss[1] {
	case "C":
		code = 1
	case "D":
		fallthrough
	case "DI":
		code = 2
	case "H":
		code = 3
	case "I":
		code = 4
	}
	offset, _ := strconv.ParseUint(ss[2], 10, 16)
	//offset, _ := strconv.Atoi(ss[2])
	return &Address{
		Code:   code,
		Offset: uint16(offset),
	}, nil
}
