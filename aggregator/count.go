package aggregator

import "context"

type countAggregator struct {
	baseAggregator
}

func (a *countAggregator) Evaluate() (val float64, err error) {
	l := len(a.targets)
	if l == 0 {
		return 0, nil
	}
	var res interface{}
	val = 0
	for _, t := range a.targets {
		res, err = a.expression(context.Background(), t)
		if err != nil {
			return
		}
		if res.(bool) {
			val++
		}
	}
	return val, nil
}
