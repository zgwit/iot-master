package types

import (
	"fmt"
	"strings"
)

type DataType int

const (
	TypeNone DataType = iota
	TypeBit
	TypeByte
	TypeWord
	TypeDWord
	TypeQWord
	TypeShort
	TypeInteger
	TypeLong
	TypeFloat
	TypeDouble
)

func (dt DataType) Parse(tp string) error {
	strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "bit":
		dt = TypeBit
	case "byte":
		dt = TypeByte
	case "word":
		fallthrough
	case "uint16":
		dt = TypeWord
	case "dword":
		fallthrough
	case "uint32":
		dt = TypeDWord
	case "qword":
		fallthrough
	case "uint64":
		dt = TypeQWord
	case "short":
		fallthrough
	case "int16":
		dt = TypeShort
	case "integer":
		fallthrough
	case "int32":
		fallthrough
	case "int":
		dt = TypeInteger
	case "long":
		fallthrough
	case "int64":
		dt = TypeLong
	case "float":
		dt = TypeFloat
	case "double":
		fallthrough
	case "float64":
		dt = TypeDouble
	default:
		return fmt.Errorf("Unknown data type: %s ", tp)
	}
	return nil
}

func (dt DataType) String() string {
	var str string
	switch dt {
	case TypeBit:
		str = "bit"
	case TypeByte:
		str = "byte"
	case TypeWord:
		str = "word"
	case TypeDWord:
		str = "dword"
	case TypeQWord:
		str = "qword"
	case TypeShort:
		str = "short"
	case TypeInteger:
		str = "integer"
	case TypeLong:
		str = "long"
	case TypeFloat:
		str = "float"
	case TypeDouble:
		str = "double"
	default:
		str = "none"
	}
	return str
}

func (dt DataType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

func (dt DataType) UnmarshalJSON(buf []byte) error {
	return dt.Parse(string(buf))
}
