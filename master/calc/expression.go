package calc

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

type Context map[string]interface{}

type Expression struct {
	program *vm.Program
}

func NewExpression(input string) (*Expression, error) {
	program, err := expr.Compile(input)
	if err != nil {
		return nil, err
	}
	return &Expression{program: program}, nil
}

func (c *Expression) Evaluate(ctx Context) (interface{}, error) {
	return expr.Run(c.program, ctx)
}
