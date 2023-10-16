package aggregator

import (
	"context"
)

type countAggregator struct {
	baseAggregator
	count float64
}

func (a *countAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalBool(context.Background(), ctx)
	if err != nil {
		return err
	}
	if res {
		a.count++
	}
	return nil
}

func (a *countAggregator) Pop() (val float64, err error) {
	//if a.count == 0 {
	//	return 0, ErrorBlank
	//}
	val = a.count
	a.count = 0
	return
}
