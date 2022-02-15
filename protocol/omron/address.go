package omron

import (
	"errors"
	"strconv"
	"strings"
)

type Command struct {
	BitCode  byte
	WordCode byte
}

type Address struct {
	Code  byte
	Addr  uint64
	Bit   uint8
	IsBit bool
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

func ParseAddress(address string) (addr Address, err error) {

	k := strings.ToUpper(address[:1])
	if cmd, ok := commands[k]; ok {

		a := address[1:]

		i := strings.IndexByte(a, '.')
		if i > -1 {
			addr.Code = cmd.BitCode
			addr.IsBit = true

			var bit int
			bit, err = strconv.Atoi(a[i+1:])
			//TODO 判断bit，只能在 0~15之间
			addr.Bit = uint8(bit)
			addr.Addr, err = strconv.ParseUint(a[:i], 10, 16)
		} else {
			addr.Code = cmd.WordCode
			addr.IsBit = false
			addr.Bit = 0
			addr.Addr, err = strconv.ParseUint(a, 10, 16)
		}
		return
	} else {
		err = errors.New("未支持通道")
	}

	return
}
