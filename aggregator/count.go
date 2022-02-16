package aggregator

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
		res, err = a.expression.Evaluate(t)
		if err != nil {
			return
		}
		if res.(bool) {
			val++
		}
	}
	return val, nil
}
