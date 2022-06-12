package aggregator

import (
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/model"
)

type Aggregator interface {
	Model() *model.Aggregator
	Init() error
	Push(ctx map[string]interface{})
	Clear()
	Evaluate() (val float64, err error)
}

//baseAggregator 聚合器
type baseAggregator struct {
	model.Aggregator

	expression gval.Evaluable
	targets    []map[string]interface{}
}

func (a *baseAggregator) Model() *model.Aggregator {
	return &a.Aggregator
}

//Init 初始化
func (a *baseAggregator) Init() (err error) {
	a.targets = make([]map[string]interface{}, 0)
	a.expression, err = calc.Language.NewEvaluable(a.Expression)
	return
}

//Push 加入
func (a *baseAggregator) Push(ctx map[string]interface{}) {
	a.targets = append(a.targets, ctx)
}

//Clear 清空
func (a *baseAggregator) Clear() {
	a.targets = make([]map[string]interface{}, 0)
}

// New 新建
func New(m *model.Aggregator) (agg Aggregator, err error) {
	switch m.Type {
	case "SUM":
		agg = &sumAggregator{baseAggregator{Aggregator: *m}}
	case "AVG":
		agg = &avgAggregator{baseAggregator{Aggregator: *m}}
	case "COUNT":
		agg = &countAggregator{baseAggregator{Aggregator: *m}}
	case "MIN":
		agg = &minAggregator{baseAggregator{Aggregator: *m}}
	case "MAX":
		agg = &maxAggregator{baseAggregator{Aggregator: *m}}
	case "FIRST":
		agg = &firstAggregator{baseAggregator{Aggregator: *m}}
	case "LAST":
		agg = &lastAggregator{baseAggregator{Aggregator: *m}}
	default:
		err = fmt.Errorf("Unknown type %s ", m.Type)
	}

	return
}
