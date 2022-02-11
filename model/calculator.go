package model

import "github.com/zgwit/iot-master/common"

type Calculator struct {
	Expression string `json:"expression"`
	Variable   string `json:"variable"`

	expr *common.Expression
	//ctx  *Context
}

func (c *Calculator) Init() (err error) {
	c.expr, err = common.NewExpression(c.Expression)
	return
}

func (c *Calculator) Evaluate(ctx common.Context) (float64, error) {
	val, err := c.expr.Evaluate(ctx)
	if err != nil {
		return 0, err
	}
	return val.(float64), nil
}
