package aggregator

import (
	"github.com/zgwit/iot-master/master/calc"
	"github.com/zgwit/iot-master/master/select"
	"math"
)

type Target struct {
	context    calc.Context
	expression *calc.Expression
}

type Aggregator struct {
	Type       Type           `json:"type"`
	As         string         `json:"as"`
	From       string         `json:"from"`
	Select     _select.Select `json:"select"`
	Expression string         `json:"expression"`

	expression *calc.Expression
	//targets []*common.Expression
	targets []calc.Context
}


func (a *Aggregator) Init() (err error) {
	a.targets = make([]calc.Context, 0)
	a.expression, err = calc.NewExpression(a.Expression)
	return
}

func (a *Aggregator) Push(ctx calc.Context) {
	a.targets = append(a.targets, ctx)
}

func (a *Aggregator) Clear() {
	a.targets = make([]calc.Context, 0)
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
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			val += res.(float64)
		}
	case COUNT:
		for _, t := range a.targets {
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			if res.(bool) {
				val++
			}
		}
	case AVG:
		for _, t := range a.targets {
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			val += res.(float64)
		}
		val = val / float64(len(a.targets))
	case MIN:
		val = math.MaxFloat64
		for _, t := range a.targets {
			res, err = a.expression.Evaluate(t)
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
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			v := res.(float64)
			if v > val {
				val = v
			}
		}
	case FIRST:
		res, err = a.expression.Evaluate(a.targets[0])
		if err != nil {
			return
		}
		val = res.(float64)
	case LAST:
		res, err = a.expression.Evaluate(a.targets[l-1])
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
