package types

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v4/pkg/calc"
)

type Filter struct {
	Field      string `json:"field"`         //字段
	Expression string `json:"expression"`    //表达式
	All        bool   `json:"all,omitempty"` //

	expression gval.Evaluable
}

func (c *Filter) Check(ctx map[string]any) (ret bool, err error) {
	if c.expression == nil {
		c.expression, err = calc.New(c.Expression)
		if err != nil {
			return
		}
	}
	ret, err = c.expression.EvalBool(context.Background(), ctx)
	return
}
