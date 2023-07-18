package aggregator

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/model"
)

type Aggregator interface {
	Compile(expression string) error
	Push(ctx map[string]interface{}) error
	Pop() (float64, error)
}

// New 新建
func New(m *model.ModAggregator, pop func(float64, error)) (agg Aggregator, err error) {
	switch m.Type {
	case "inc", "increase":
		agg = &incAggregator{}
	case "dec", "decrease":
		agg = &decAggregator{}
	case "sum":
		agg = &sumAggregator{}
	case "avg", "average":
		agg = &avgAggregator{}
	case "count":
		agg = &countAggregator{}
	case "min", "minimum":
		agg = &minAggregator{}
	case "max", "maximum":
		agg = &maxAggregator{}
	case "first":
		agg = &firstAggregator{}
	case "last":
		agg = &lastAggregator{}
	default:
		err = fmt.Errorf("Unknown type %s ", m.Type)
		return
	}
	//agg.expression, err = calc.New(a.Expression)
	err = agg.Compile(m.Expression)

	_, err = _cron.AddFunc(m.Crontab, func() {
		//并发处理，如果设备太多就完蛋了
		go pop(agg.Pop())
	})

	if err != nil {
		return nil, err
	}

	return
}
