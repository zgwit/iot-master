package aggregator

type avgAggregator struct {
	baseAggregator
}

func (a *avgAggregator) Evaluate() (val float64, err error) {
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
		val += res.(float64)
	}
	val = val / float64(len(a.targets))
	return val, nil
}
