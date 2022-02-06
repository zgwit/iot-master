package interval

import "github.com/asaskevich/EventBus"

type Calculator struct {
	Expression string `json:"expression"`
	Variable   string `json:"variable"`

	expr *Expression
	//ctx  *Context

	events EventBus.Bus
}

func (c *Calculator) Init(ctx Context) error {
	//事件队列
	c.events = EventBus.New()

	var err error
	c.expr, err = NewExpression(c.Expression, ctx)
	if err != nil {
		return err
	}
	return nil
}

//TODO 应该放到device中直接执行
func (c *Calculator) Evaluate() error {
	val, err := c.expr.Evaluate()
	if err != nil {
		return err
	}
	//回写
	//(*c.expr.ctx)[c.Variable] = val
	c.events.Publish("data", Context{c.Variable: val}) //使用事件回传
	return nil
}
