package interval

type Calculator struct {
	Expression string `json:"expression"`
	Variable   string `json:"variable"`

	expr *Expression
	//ctx  *Context
}

func (c *Calculator) Compile(ctx  *Context) error {

	var err error
	c.expr, err = NewExpression(c.Expression, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Calculator) Evaluate() error {
	val, err := c.expr.Evaluate()
	if err != nil {
		return err
	}
	//回写
	(*c.expr.ctx)[c.Variable] = val
	return nil
}
