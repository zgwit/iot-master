package aggregator

type lastAggregator struct {
	baseAggregator
}

func (a *lastAggregator) Evaluate() (val float64, err error) {
	l := len(a.targets)
	if l == 0 {
		return 0, nil
	}
	var res interface{}
	res, err = a.expression.Evaluate(a.targets[l-1])
	if err != nil {
		return
	}
	val = res.(float64)
	return val, nil
}
