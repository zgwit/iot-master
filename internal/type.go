package interval

import (
	"fmt"
	"strings"
)

type DataType int

const (
	TypeNONE DataType = iota
	TypeBIT
	TypeBYTE
	TypeWORD
	TypeDWORD
	TypeQWORD
	TypeSHORT
	TypeINTEGER
	TypeLONG
	TypeFLOAT
	TypeDOUBLE
)

func (dt DataType) Parse(tp string) error {
	strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "bit":
		dt = TypeBIT
	case "byte":
		dt = TypeBYTE
	case "word":
		fallthrough
	case "uint16":
		dt = TypeWORD
	case "dword":
		fallthrough
	case "uint32":
		dt = TypeDWORD
	case "qword":
		fallthrough
	case "uint64":
		dt = TypeQWORD
	case "short":
		fallthrough
	case "int16":
		dt = TypeSHORT
	case "integer":
		fallthrough
	case "int32":
		fallthrough
	case "int":
		dt = TypeINTEGER
	case "long":
		fallthrough
	case "int64":
		dt = TypeLONG
	case "float":
		dt = TypeFLOAT
	case "double":
		fallthrough
	case "float64":
		dt = TypeDOUBLE
	default:
		return fmt.Errorf("Unknown data type: %s ", tp)
	}
	return nil
}

func (dt DataType) String() string {
	var str string
	switch dt {
	case TypeBIT:
		str = "bit"
	case TypeBYTE:
		str = "byte"
	case TypeWORD:
		str = "word"
	case TypeDWORD:
		str = "dword"
	case TypeQWORD:
		str = "qword"
	case TypeSHORT:
		str = "short"
	case TypeINTEGER:
		str = "integer"
	case TypeLONG:
		str = "long"
	case TypeFLOAT:
		str = "float"
	case TypeDOUBLE:
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
