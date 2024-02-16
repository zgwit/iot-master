package types

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v4/pkg/calc"
)

type Calculator struct {
	Assign     string `json:"assign"`     //赋值
	Expression string `json:"expression"` //表达式
	expression gval.Evaluable
}

func (c *Calculator) Eval(ctx map[string]any) (err error) {
	if c.expression == nil {
		c.expression, err = calc.New(c.Expression)
		if err != nil {
			return
		}
	}
	ctx[c.Assign], err = c.expression(context.Background(), ctx)
	return
}
