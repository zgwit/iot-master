package interval

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

type Context map[string]interface{}

type Expression struct {
	program *vm.Program
	ctx  *Context
}

func NewExpression(input string, ctx *Context) (*Expression, error) {
	program, err := expr.Compile(input, expr.Env(*ctx))
	if err != nil {
		return nil, err
	}
	return &Expression{program: program, ctx: ctx}, nil
}

func (c *Expression) Evaluate() (interface{}, error) {
	return c.Evaluate()
}



