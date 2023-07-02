package aggregator

import (
	"context"
)

type decAggregator struct {
	baseAggregator

	hasLast bool

	last    float64
	current float64

	dirty bool
}

func (a *decAggregator) Push(ctx map[string]interface{}) error {
	res, err := a.expression.EvalFloat64(context.Background(), ctx)
	if err != nil {
		return err
	}
	a.current = res
	a.dirty = true

	return nil
}

func (a *decAggregator) Pop() (val float64, err error) {
	if !a.dirty {
		return 0, ErrorBlank
	}
	if !a.hasLast {
		a.hasLast = true
		a.last = a.current
		a.dirty = false
		return 0, ErrorBlank
	}

	val = a.current - a.last
	a.last = a.current
	a.dirty = false

	return
}
