package aggregator

import (
	"fmt"
)

type Aggregator struct {
	Crontab    string `json:"crontab"`    //定时计划
	Expression string `json:"expression"` //表达式
	Type       string `json:"type"`       //聚合算法 inc dec avg count min max sum last first
	Assign     string `json:"assign"`     //赋值
}

type Instance interface {
	Compile(expression string) error
	Push(ctx map[string]interface{}) error
	Pop() (float64, error)
}

// New 新建
func New(m *Aggregator, pop func(float64, error)) (agg Instance, err error) {
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
