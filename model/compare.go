package model

import (
	"fmt"
	"strings"
)

//CompareType 数据类型
type CompareType int

const (
	//CompareNONE 空类型
	CompareNONE CompareType = iota
	CompareEQ
	CompareNE
	CompareLT
	CompareLE
	CompareGT
	CompareGE
	CompareBetween
	CompareOver
)

func (dt *CompareType) Eval(value, value1, value2 float64) bool {
	var ret bool
	switch *dt {
	case CompareNONE:
		ret = false
	case CompareEQ:
		ret = value == value1
	case CompareNE:
		ret = value != value1
	case CompareLT:
		ret = value < value1
	case CompareLE:
		ret = value <= value1
	case CompareGT:
		ret = value > value1
	case CompareGE:
		ret = value >= value1
	case CompareBetween:
		ret = value >= value1 && value <= value2
	case CompareOver:
		ret = value < value1 || value > value2
	default:
		ret = false
	}
	return ret
}

//Parse 解析类型
func (dt *CompareType) Parse(tp string) error {
	//var *dt DataType
	tp = tp[1 : len(tp)-1]
	//strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "none":
		*dt = CompareNONE
	case "=":
		*dt = CompareEQ
	case "!=":
		*dt = CompareNE
	case "<":
		*dt = CompareLT
	case "<=":
		*dt = CompareLE
	case ">":
		*dt = CompareGT
	case ">=":
		*dt = CompareGE
	case "><":
		*dt = CompareBetween
	case "<>":
		*dt = CompareOver
	default:
		return fmt.Errorf("Unknown data type: %s ", tp)
	}
	return nil
}

//String 转化成字符串
func (dt *CompareType) String() string {
	var str string
	switch *dt {
	case CompareNONE:
		str = "none"
	case CompareEQ:
		str = "="
	case CompareNE:
		str = "!="
	case CompareLT:
		str = "<"
	case CompareLE:
		str = "<="
	case CompareGT:
		str = ">"
	case CompareGE:
		str = ">="
	case CompareBetween:
		str = "><"
	case CompareOver:
		str = "<>"
	default:
		str = "none"
	}
	return str
}

//MarshalJSON 序列化
func (dt *CompareType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

//UnmarshalJSON 解析
func (dt *CompareType) UnmarshalJSON(buf []byte) error {
	return dt.Parse(string(buf))
}
