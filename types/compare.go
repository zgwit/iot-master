package types

import (
	"fmt"
	"strings"
)

type CompareType int

const (
	CompareNone CompareType = iota
	CompareLT
	CompareLE
	CompareEQ
	CompareNE
	CompareGT
	CompareGE
)

func (ct CompareType) Parse(tp string) error {
	strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "lt":
		ct = CompareLT
	case "le":
		ct = CompareLE
	case "eq":
		ct = CompareEQ
	case "ne":
		ct = CompareNE
	case "gt":
		ct = CompareGT
	case "ge":
		ct = CompareGE
	default:
		return fmt.Errorf("Unknown compare type: %s ", tp)
	}
	return nil
}

func (ct CompareType) String() string {
	var str string
	switch ct {
	case CompareLT:
		str = "lt"
	case CompareLE:
		str = "le"
	case CompareEQ:
		str = "eq"
	case CompareNE:
		str = "ne"
	case CompareGT:
		str = "gt"
	case CompareGE:
		str = "ge"
	default:
		str = "none"
	}
	return str
}

func (ct CompareType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.String() + `"`), nil
}

func (ct CompareType) UnmarshalJSON(buf []byte) error {
	return ct.Parse(string(buf))
}

type Compare struct {
	Type     CompareType `json:"type"`
	Value    float64     `json:"value"`
	Variable *Variable   `json:"variable"`
}

func (comp *Compare) Evaluate() bool {
	switch comp.Type {
	case CompareLT:
		return comp.Variable.Value < comp.Value
	case CompareLE:
		return comp.Variable.Value <= comp.Value
	case CompareEQ:
		return comp.Variable.Value == comp.Value
	case CompareNE:
		return comp.Variable.Value != comp.Value
	case CompareGT:
		return comp.Variable.Value > comp.Value
	case CompareGE:
		return comp.Variable.Value >= comp.Value
	}
	return false
}
