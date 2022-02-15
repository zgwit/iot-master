package calc

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

//Context 上下文
type Context map[string]interface{}

//Expression 表达式
type Expression struct {
	program *vm.Program
}

//NewExpression 编译
func NewExpression(input string) (*Expression, error) {
	program, err := expr.Compile(input)
	if err != nil {
		return nil, err
	}
	return &Expression{program: program}, nil
}

//Evaluate 计算
func (c *Expression) Evaluate(ctx Context) (interface{}, error) {
	return expr.Run(c.program, ctx)
}
