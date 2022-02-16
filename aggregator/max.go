package aggregator

import "math"

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
		res, err = a.expression.Evaluate(t)
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
