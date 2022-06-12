package aggregator

import "context"

type firstAggregator struct {
	baseAggregator
}

func (a *firstAggregator) Evaluate() (val float64, err error) {
	l := len(a.targets)
	if l == 0 {
		return 0, nil
	}
	var res interface{}
	res, err = a.expression(context.Background(), a.targets[0])
	if err != nil {
		return
	}
	val = res.(float64)
	return val, nil
}
