package aggregator

import (
	"fmt"
	"strings"
)

type Type int

const (
	NONE Type = iota
	SUM
	COUNT
	AVG
	MIN
	MAX
	FIRST
	LAST
)

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

func (t Type) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t Type) UnmarshalJSON(buf []byte) error {
	return t.Parse(string(buf))
}

