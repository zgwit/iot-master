package interval

import (
	"fmt"
	"math"
	"strings"
)

type AggregatorType int

const (
	AggregatorNONE AggregatorType = iota
	AggregatorSUM
	AggregatorCOUNT
	AggregatorAVG
	AggregatorMIN
	AggregatorMAX
	AggregatorFIRST
	AggregatorLAST
)

func (ct AggregatorType) Parse(tp string) error {
	strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "sum":
		ct = AggregatorSUM
	case "count":
		ct = AggregatorCOUNT
	case "avg":
		ct = AggregatorAVG
	case "min":
		ct = AggregatorMIN
	case "max":
		ct = AggregatorMAX
	case "first":
		ct = AggregatorFIRST
	case "last":
		ct = AggregatorLAST
	default:
		return fmt.Errorf("Unknown compare type: %s ", tp)
	}
	return nil
}

func (ct AggregatorType) String() string {
	var str string
	switch ct {
	case AggregatorSUM:
		str = "sum"
	case AggregatorCOUNT:
		str = "count"
	case AggregatorAVG:
		str = "avg"
	case AggregatorMIN:
		str = "min"
	case AggregatorMAX:
		str = "max"
	case AggregatorFIRST:
		str = "first"
	case AggregatorLAST:
		str = "last"
	default:
		str = "none"
	}
	return str
}

func (ct AggregatorType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.String() + `"`), nil
}

func (ct AggregatorType) UnmarshalJSON(buf []byte) error {
	return ct.Parse(string(buf))
}

type AggregatorTarget struct {
	device     *Device
	expression *Expression
}

type Aggregator struct {
	Type       AggregatorType `json:"type"`
	As         string         `json:"as"`
	From       string         `json:"from"`
	Expression string         `json:"expression"`

	Tags []string `json:"tags"`

	targets []AggregatorTarget
	ctx     *Context
}

func hasTag(a, b []string) bool {
	for i := len(a); i >= 0; i-- {
		for j := len(b); j >= 0; j-- {
			//if a[i] == b[j] {
			if strings.EqualFold(a[i], b[j]) {
				return true
			}
		}
	}
	return false
}

func (a *Aggregator) Compile(devices []*Device, ctx *Context) error {
	a.ctx = ctx
	for _, dev := range devices {
		if hasTag(a.Tags, dev.Tags) {
			expr, err := NewExpression(a.Expression, &dev.Context)
			if err != nil {
				return err
			}
			a.targets = append(a.targets, AggregatorTarget{device: dev, expression: expr})
		}
	}
	return nil
}

func (a *Aggregator) Evaluate() error {
	l := len(a.targets)
	if l == 0 {
		return nil
	}

	var ret float64 = 0
	switch a.Type {
	case AggregatorSUM:
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			ret += val.(float64)
		}
	case AggregatorCOUNT:
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			if val.(bool) {
				ret++
			}
		}
	case AggregatorAVG:
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			ret += val.(float64)
		}
		ret = ret / float64(len(a.targets))
	case AggregatorMIN:
		ret = math.MaxFloat64
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			v := val.(float64)
			if v < ret {
				ret = v
			}
		}
	case AggregatorMAX:
		ret = math.SmallestNonzeroFloat32
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			v := val.(float64)
			if v > ret {
				ret = v
			}
		}
	case AggregatorFIRST:
		val, err := a.targets[0].expression.Evaluate()
		if err != nil {
			return err
		}
		ret = val.(float64)
	case AggregatorLAST:
		val, err := a.targets[l-1].expression.Evaluate()
		if err != nil {
			return err
		}
		ret = val.(float64)
	default:
		return nil //TODO error
	}

	//写入结果
	(*a.ctx)[a.As] = ret

	return nil
}
