package aggregator

import (
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/model"
	"math"
)

//Target 目标
type Target struct {
	context    calc.Context
	expression *calc.Expression
}

//Aggregator 聚合器
type Aggregator struct {
	model.Aggregator

	expression *calc.Expression
	//targets []*common.Expression
	targets []calc.Context
}

//Init 初始化
func (a *Aggregator) Init() (err error) {
	a.targets = make([]calc.Context, 0)
	a.expression, err = calc.NewExpression(a.Expression)
	return
}

//Push 加入
func (a *Aggregator) Push(ctx calc.Context) {
	a.targets = append(a.targets, ctx)
}

//Clear 清空
func (a *Aggregator) Clear() {
	a.targets = make([]calc.Context, 0)
}

//Evaluate 计算
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
	case "SUM":
		for _, t := range a.targets {
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			val += res.(float64)
		}
	case "COUNT":
		for _, t := range a.targets {
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			if res.(bool) {
				val++
			}
		}
	case "AVG":
		for _, t := range a.targets {
			res, err = a.expression.Evaluate(t)
			if err != nil {
				return
			}
			val += res.(float64)
		}
		val = val / float64(len(a.targets))
	case "MIN":
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
	case "MAX":
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
	case "FIRST":
		res, err = a.expression.Evaluate(a.targets[0])
		if err != nil {
			return
		}
		val = res.(float64)
	case "LAST":
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
