package aggregator

import (
	"fmt"
	"strings"
)

//Type 类型
type Type int

const (
	//NONE 空类型
	NONE Type = iota
	SUM
	COUNT
	AVG
	MIN
	MAX
	FIRST
	LAST
)

//Parse 解析
func (t Type) Parse(tp string) error {
	strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "sum":
		t = SUM
	case "count":
		t = COUNT
	case "avg":
		t = AVG
	case "min":
		t = MIN
	case "max":
		t = MAX
	case "first":
		t = FIRST
	case "last":
		t = LAST
	default:
		return fmt.Errorf("Unknown compare type: %s ", tp)
	}
	return nil
}

//String 字符串
func (t Type) String() string {
	var str string
	switch t {
	case SUM:
		str = "sum"
	case COUNT:
		str = "count"
	case AVG:
		str = "avg"
	case MIN:
		str = "min"
	case MAX:
		str = "max"
	case FIRST:
		str = "first"
	case LAST:
		str = "last"
	default:
		str = "none"
	}
	return str
}

//MarshalJSON 序列化
func (t Type) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

//UnmarshalJSON 反序列化
func (t Type) UnmarshalJSON(buf []byte) error {
	return t.Parse(string(buf))
}
