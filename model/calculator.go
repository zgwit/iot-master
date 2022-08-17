package model

import (
	"context"
	"github.com/PaesslerAG/gval"
	"iot-master/pkg/calc"
)

//Calculator 计算器
type Calculator struct {
	Expression string `json:"expression"`
	As         string `json:"as"`
	Store      bool   `json:"store"`

	//ctx  *Context
	expr gval.Evaluable
}

//Init 初始化（编译）
func (c *Calculator) Init() (err error) {
	c.expr, err = calc.Language.NewEvaluable(c.Expression)
	return
}

//Evaluate 计算
func (c *Calculator) Evaluate(ctx map[string]interface{}) (interface{}, error) {
	return c.expr(context.Background(), ctx)
}
