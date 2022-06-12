package aggregator

import (
	"context"
	"math"
)

type maxAggregator struct {
	baseAggregator
}

func (a *maxAggregator) Evaluate() (val float64, err error) {
	l := len(a.targets)
	if l == 0 {
		return 0, nil
	}
	var res interface{}
	val = math.SmallestNonzeroFloat32
	for _, t := range a.targets {
		res, err = a.expression(context.Background(), t)
		if err != nil {
			return
		}
		v := res.(float64)
		if v > val {
			val = v
		}
	}
	return val, nil
}
