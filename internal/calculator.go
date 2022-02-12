package internal

import "github.com/zgwit/iot-master/internal/calc"

type Calculator struct {
	Expression string `json:"expression"`
	Variable   string `json:"variable"`

	expr *calc.Expression
	//ctx  *Context
}

func (c *Calculator) Init() (err error) {
	c.expr, err = calc.NewExpression(c.Expression)
	return
}

func (c *Calculator) Evaluate(ctx calc.Context) (float64, error) {
	val, err := c.expr.Evaluate(ctx)
	if err != nil {
		return 0, err
	}
	return val.(float64), nil
}
