package model

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/protocols/helper"
	"math"
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
func (dt *DataType) Encode(value interface{}, le bool, precision int) []byte {
	buf := make([]byte, 8)
	switch *dt {
	case TypeBIT:
		if helper.ToBool(value) {
			buf[0] = 1 //?????
		} else {
			buf[0] = 0
		}
	case TypeBYTE:
		buf[0] = helper.ToUint8(value)
	case TypeWORD:
		var val uint16
		if precision > 0 {
			val = uint16(helper.ToFloat64(value) * math.Pow10(precision))
		} else {
			val = helper.ToUint16(value)
		}
		if le {
			helper.WriteUint16LittleEndian(buf, val)
		} else {
			helper.WriteUint16(buf, val)
		}
	case TypeDWORD:
		var val uint32
		if precision > 0 {
			val = uint32(helper.ToFloat64(value) * math.Pow10(precision))
		} else {
			val = helper.ToUint32(value)
		}
		if le {
			helper.WriteUint32LittleEndian(buf, val)
		} else {
			helper.WriteUint32(buf, val)
		}
	case TypeQWORD:
		var val uint64
		if precision > 0 {
			val = uint64(value.(float64) * math.Pow10(precision))
		} else {
			val = helper.ToUint64(value)
		}
		if le {
			helper.WriteUint64LittleEndian(buf, val)
		} else {
			helper.WriteUint64(buf, val)
		}
	case TypeSHORT:
		var val int16
		if precision > 0 {
			val = int16(helper.ToFloat64(value) * math.Pow10(precision))
		} else {
			val = helper.ToInt16(value)
		}
		if le {
			helper.WriteUint16LittleEndian(buf, uint16(val))
		} else {
			helper.WriteUint16(buf, uint16(val))
		}
	case TypeINTEGER:
		var val int32
		if precision > 0 {
			val = int32(helper.ToFloat64(value) * math.Pow10(precision))
		} else {
			val = helper.ToInt32(value)
		}
		if le {
			helper.WriteUint32LittleEndian(buf, uint32(val))
		} else {
			helper.WriteUint32(buf, uint32(val))
		}
	case TypeLONG:
		var val int64
		if precision > 0 {
			val = int64(helper.ToFloat64(value) * math.Pow10(precision))
		} else {
			val = helper.ToInt64(value)
		}
		if le {
			helper.WriteUint64LittleEndian(buf, uint64(val))
		} else {
			helper.WriteUint64(buf, uint64(val))
		}
	case TypeFLOAT:
		val := helper.ToFloat32(value)
		if le {
			helper.WriteFloat32LittleEndian(buf, val)
		} else {
			helper.WriteFloat32(buf, val)
		}
	case TypeDOUBLE:
		val := helper.ToFloat64(value)
		if le {
			helper.WriteFloat64LittleEndian(buf, val)
		} else {
			helper.WriteFloat64(buf, val)
		}
	default:
		//TODO error
	}
	return buf[:dt.Size()]
}

//Decode 解码
func (dt *DataType) Decode(buf []byte, le bool, precision int) (val interface{}, err error) {
	switch *dt {
	case TypeBIT:
		if buf[0] > 0 {
			val = true
		} else {
			val = false
		}
	case TypeBYTE:
		val = buf[0]
	case TypeWORD:
		var value uint16
		if le {
			value = helper.ParseUint16LittleEndian(buf)
		} else {
			value = helper.ParseUint16(buf)
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		}
	case TypeDWORD:
		var value uint32
		if le {
			value = helper.ParseUint32LittleEndian(buf)
		} else {
			value = helper.ParseUint32(buf)
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		}
	case TypeQWORD:
		var value uint64
		if le {
			value = helper.ParseUint64LittleEndian(buf)
		} else {
			value = helper.ParseUint64(buf)
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		}
	case TypeSHORT:
		var value int16
		if le {
			value = int16(helper.ParseUint16LittleEndian(buf))
		} else {
			value = int16(helper.ParseUint16(buf))
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		}
	case TypeINTEGER:
		var value int32
		if le {
			value = int32(helper.ParseUint32LittleEndian(buf))
		} else {
			value = int32(helper.ParseUint32(buf))
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		}
	case TypeLONG:
		var value int64
		if le {
			value = int64(helper.ParseUint64LittleEndian(buf))
		} else {
			value = int64(helper.ParseUint64(buf))
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		}
	case TypeFLOAT:
		if le {
			val = helper.ParseFloat32LittleEndian(buf)
		} else {
			val = helper.ParseFloat32(buf)
		}
	case TypeDOUBLE:
		if le {
			val = helper.ParseFloat64LittleEndian(buf)
		} else {
			val = helper.ParseFloat64(buf)
		}
	default:
		err = errors.New("未知的数据类型")
	}
	return
}

//MarshalJSON 序列化
func (dt *DataType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

//UnmarshalJSON 解析
func (dt *DataType) UnmarshalJSON(buf []byte) error {
	return dt.Parse(string(buf))
}
