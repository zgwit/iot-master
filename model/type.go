package model

import (
	"fmt"
	"github.com/zgwit/iot-master/protocol/helper"
	"strings"
)

//DataType 数据类型
type DataType int

const (
	//TypeNONE 空类型
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

//Parse 解析类型
func (dt *DataType) Parse(tp string) error {
	//var *dt DataType
	tp = tp[1 : len(tp)-1]
	//strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "none":
		*dt = TypeNONE
	case "bit":
		*dt = TypeBIT
	case "byte":
		*dt = TypeBYTE
	case "word":
		fallthrough
	case "uint16":
		*dt = TypeWORD
	case "dword":
		fallthrough
	case "uint32":
		*dt = TypeDWORD
	case "qword":
		fallthrough
	case "uint64":
		*dt = TypeQWORD
	case "short":
		fallthrough
	case "int16":
		*dt = TypeSHORT
	case "integer":
		fallthrough
	case "int32":
		fallthrough
	case "int":
		*dt = TypeINTEGER
	case "long":
		fallthrough
	case "int64":
		*dt = TypeLONG
	case "float":
		*dt = TypeFLOAT
	case "double":
		fallthrough
	case "float64":
		*dt = TypeDOUBLE
	default:
		return fmt.Errorf("Unknown data type: %s ", tp)
	}
	return nil
}

//String 转化成字符串
func (dt *DataType) String() string {
	var str string
	switch *dt {
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

//Size 宽度
func (dt *DataType) Size() int {
	var s int
	switch *dt {
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

//Encode 编码
func (dt *DataType) Encode(val float64, le bool) []byte {
	buf := make([]byte, 8)
	switch *dt {
	case TypeBIT:
		if val > 0 {
			buf[0] = 1 //?????
		} else {
			buf[0] = 0
		}
	case TypeBYTE:
		buf[0] = uint8(val)
	case TypeWORD:
		if le {
			helper.WriteUint16LittleEndian(buf, uint16(val))
		} else {
			helper.WriteUint16(buf, uint16(val))
		}
	case TypeDWORD:
		if le {
			helper.WriteUint32LittleEndian(buf, uint32(val))
		} else {
			helper.WriteUint32(buf, uint32(val))
		}
	case TypeQWORD:
		if le {
			helper.WriteUint64LittleEndian(buf, uint64(val))
		} else {
			helper.WriteUint64(buf, uint64(val))
		}
	case TypeSHORT:
		if le {
			helper.WriteUint16LittleEndian(buf, uint16(int16(val)))
		} else {
			helper.WriteUint16(buf, uint16(int16(val)))
		}
	case TypeINTEGER:
		if le {
			helper.WriteUint32LittleEndian(buf, uint32(int32(val)))
		} else {
			helper.WriteUint32(buf, uint32(int32(val)))
		}
	case TypeLONG:
		if le {
			helper.WriteUint64LittleEndian(buf, uint64(int64(val)))
		} else {
			helper.WriteUint64(buf, uint64(int64(val)))
		}
	case TypeFLOAT:
		if le {
			helper.WriteFloat32LittleEndian(buf, float32(val))
		} else {
			helper.WriteFloat32(buf, float32(val))
		}
	case TypeDOUBLE:
		if le {
			helper.WriteFloat64LittleEndian(buf, val)
		} else {
			helper.WriteFloat64(buf, val)
		}
	default:
		//TODO error
	}
	return nil
}

//Decode 解码
func (dt *DataType) Decode(buf []byte, le bool) (float64, error) {
	var val float64
	switch *dt {
	case TypeBIT:
		if buf[0] > 0 {
			val = 1
		} else {
			val = 0
		}
	case TypeBYTE:
		val = float64(buf[0])
	case TypeWORD:
		if le {
			val = float64(helper.ParseUint16LittleEndian(buf))
		} else {
			val = float64(helper.ParseUint16(buf))
		}
	case TypeDWORD:
		if le {
			val = float64(helper.ParseUint32LittleEndian(buf))
		} else {
			val = float64(helper.ParseUint32(buf))
		}
	case TypeQWORD:
		if le {
			val = float64(helper.ParseUint64LittleEndian(buf))
		} else {
			val = float64(helper.ParseUint64(buf))
		}
	case TypeSHORT:
		if le {
			val = float64(int16(helper.ParseUint16LittleEndian(buf)))
		} else {
			val = float64(int16(helper.ParseUint16(buf)))
		}
	case TypeINTEGER:
		if le {
			val = float64(int32(helper.ParseUint32LittleEndian(buf)))
		} else {
			val = float64(int32(helper.ParseUint32(buf)))
		}
	case TypeLONG:
		if le {
			val = float64(int64(helper.ParseUint64LittleEndian(buf)))
		} else {
			val = float64(int64(helper.ParseUint64(buf)))
		}
	case TypeFLOAT:
		if le {
			val = float64(helper.ParseFloat32LittleEndian(buf))
		} else {
			val = float64(helper.ParseFloat32(buf))
		}
	case TypeDOUBLE:
		if le {
			val = helper.ParseFloat64LittleEndian(buf)
		} else {
			val = helper.ParseFloat64(buf)
		}
	default:
		//TODO error
	}
	return val, nil
}

//MarshalJSON 序列化
func (dt *DataType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

//UnmarshalJSON 解析
func (dt *DataType) UnmarshalJSON(buf []byte) error {
	return dt.Parse(string(buf))
}
