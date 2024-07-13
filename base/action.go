package base

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/god-jason/bucket/pkg/calc"
	"strings"
	"time"
)

type Action struct {
	ProductId  string         `json:"product_id,omitempty" bson:"product_id"`
	DeviceId   string         `json:"device_id,omitempty" bson:"device_id"`
	Name       string         `json:"name,omitempty"`
	Parameters map[string]any `json:"parameters,omitempty"`
	Delay      time.Duration  `json:"delay,omitempty"` //延迟 ms

	_parameters map[string]gval.Evaluable
}

func (a *Action) Init() (err error) {
	a._parameters = make(map[string]gval.Evaluable)
	for key, value := range a.Parameters {
		if val, ok := value.(string); ok {
			if expr, has := strings.CutPrefix(val, "="); has {
				a._parameters[key], err = calc.New(expr)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *Action) Evaluate(args any) (values map[string]any, err error) {
	values = make(map[string]any)
	for key, value := range a.Parameters {
		values[key] = value
	}
	for key, value := range a._parameters {
		if value != nil {
			values[key], err = value(context.Background(), args)
			if err != nil {
				return
			}
		}
	}
	return
}
