package interval

import "github.com/Knetic/govaluate"

type Calculator struct {
	Variable   string `json:"variable"`
	Expression string `json:"expression"`

	variable   *Variable
	expression *govaluate.EvaluableExpression
}

func (c Calculator) Evaluate() error {

	return nil
}
