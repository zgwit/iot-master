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

func (a *Aggregator) Evaluate() (val float64, err error) {
	l := len(a.targets)
	if l == 0 {
		return 0, nil
	}
	var res interface{}

	//TODO 拆成子类，多态实现？？？
	//var val float64 = 0
	val = 0
	switch a.Type {
	case SUM:
		for _, t := range a.targets {
			res, err = t.expression.Evaluate()
			if err != nil {
				return
			}
			val += res.(float64)
		}
	case COUNT:
		for _, t := range a.targets {
			res, err = t.expression.Evaluate()
			if err != nil {
				return
			}
			if res.(bool) {
				val++
			}
		}
	case AVG:
		for _, t := range a.targets {
			res, err = t.expression.Evaluate()
			if err != nil {
				return
			}
			val += res.(float64)
		}
		val = val / float64(len(a.targets))
	case MIN:
		val = math.MaxFloat64
		for _, t := range a.targets {
			res, err = t.expression.Evaluate()
			if err != nil {
				return
			}
			v := res.(float64)
			if v < val {
				val = v
			}
		}
	case MAX:
		val = math.SmallestNonzeroFloat32
		for _, t := range a.targets {
			res, err = t.expression.Evaluate()
			if err != nil {
				return
			}
			v := res.(float64)
			if v > val {
				val = v
			}
		}
	case FIRST:
		res, err = a.targets[0].expression.Evaluate()
		if err != nil {
			return
		}
		val = res.(float64)
	case LAST:
		res, err = a.targets[l-1].expression.Evaluate()
		if err != nil {
			return
		}
		val = res.(float64)
	default:
		return 0, nil //TODO error
	}

	//写入结果
	//(*a.ctx)[a.As] = val

	return
}
