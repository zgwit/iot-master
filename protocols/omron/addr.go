package omron

import (
	"errors"
	"github.com/zgwit/iot-master/protocol"
	"regexp"
	"strconv"
)

type Command struct {
	BitCode  byte
	WordCode byte
}

var commands = map[string]Command{
	//DM Area
	"D": {0x02, 0x82},
	//CIO Area
	"C": {0x30, 0xB0},
	//Work Area
	"W": {0x31, 0xB1},
	//Holding Bit Area
	"H": {0x32, 0xB2},
	//Auxiliary Bit Area
	"A": {0x33, 0xB3},
}

type Address struct {
	Code   byte
	Offset uint16
	Bits   uint8
	IsBit  bool
}

func (a *Address) String() string {
	code := ""
	switch a.Code {
	case 0x02, 0x82:
		code = "D"
	case 0x30, 0xB0:
		code = "C"
	case 0x31, 0xB1:
		code = "W"
	case 0x32, 0xB2:
		code = "H"
	case 0x33, 0xB3:
		code = "A"
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
	addrRegexp = regexp.MustCompile(`^(D|C|W|H|A)(\d+)(.\d+)?$`)
}

func ParseAddress(addr string) (protocol.Addr, error) {
	ss := addrRegexp.FindStringSubmatch(addr)
	if ss == nil || len(ss) < 3 {
		return nil, errors.New("unknown address")
	}
	var code uint8 = 1
	switch ss[1] {
	case "D":
		code = 0x82
	case "C":
		code = 0xB0
	case "W":
		code = 0xB1
	case "H":
		code = 0xB2
	case "A":
		code = 0xB3
	}
	offset, _ := strconv.ParseUint(ss[2], 10, 16)
	//offset, _ := strconv.Atoi(ss[2])
	address := &Address{
		Code:   code,
		Offset: uint16(offset),
	}
	if len(ss) == 4 {
		//TODO 判断bit，只能在 0~15之间
		bits, _ :=  strconv.ParseUint(ss[3][1:], 10, 8)
		address.Bits = uint8(bits)
		address.IsBit = true
		address.Code -= 0x80
	}
	return address, nil
}
