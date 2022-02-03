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


func (dt DataType) Size() int {
	var s int
	switch dt {
	case TypeBIT:
		s = 1
	case TypeBYTE:
		s = 1
	case TypeWORD:
		s = 2
	case TypeDWORD:
		s = 4
	case TypeQWORD:
		s = 8
	case TypeSHORT:
		s = 2
	case TypeINTEGER:
		s = 4
	case TypeLONG:
		s = 8
	case TypeFLOAT:
		s = 4
	case TypeDOUBLE:
		s = 8
	default:
		s = 1
	}
	return s
}

func (dt DataType) Encode(val float64) []byte {

	return nil
}

func (dt DataType) Decode(val []byte) (float64, error) {

	return 0, nil
}

func (dt DataType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

func (dt DataType) UnmarshalJSON(buf []byte) error {
	return dt.Parse(string(buf))
}
