package aggregator

import (
	"github.com/zgwit/iot-master/internal"
	"math"
	"strings"
)

type Target struct {
	device     *interval.Device
	expression *interval.Expression
}

type Aggregator struct {
	Type       Type   `json:"type"`
	As         string `json:"as"`
	From       string `json:"from"`
	Expression string `json:"expression"`

	Tags []string `json:"tags,omitempty"`

	targets []Target
	ctx     *interval.Context
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

func (a *Aggregator) Compile(devices []*interval.Device, ctx *interval.Context) error {
	a.ctx = ctx
	for _, dev := range devices {
		if hasTag(a.Tags, dev.Tags) {
			expr, err := interval.NewExpression(a.Expression, dev.Context)
			if err != nil {
				return err
			}
			a.targets = append(a.targets, Target{device: dev, expression: expr})
		}
	}
	return nil
}

func (a *Aggregator) Evaluate() error {
	l := len(a.targets)
	if l == 0 {
		return nil
	}

	//TODO 拆成子类，多态实现？？？
	var ret float64 = 0
	switch a.Type {
	case SUM:
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			ret += val.(float64)
		}
	case COUNT:
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			if val.(bool) {
				ret++
			}
		}
	case AVG:
		for _, t := range a.targets {
			val, err := t.expression.Evaluate()
			if err != nil {
				return err
			}
			ret += val.(float64)
		}
		ret = ret / float64(len(a.targets))
	case MIN:
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
	case MAX:
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
	case FIRST:
		val, err := a.targets[0].expression.Evaluate()
		if err != nil {
			return err
		}
		ret = val.(float64)
	case LAST:
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
