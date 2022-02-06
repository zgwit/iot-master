package interval


type Calculator struct {
	Expression string `json:"expression"`
	Variable   string `json:"variable"`

	expr *Expression
	//ctx  *Context
}

func (c *Calculator) Init(ctx Context) (err error) {
	c.expr, err = NewExpression(c.Expression, ctx)
	return
}

func (c *Calculator) Evaluate() (float64, error) {
	val, err := c.expr.Evaluate()
	if err != nil {
		return 0, err
	}
	return val.(float64), nil
}
