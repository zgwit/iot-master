package types

type Calculator struct {
	Variable   string
	Expression string
}

func (c Calculator) Evaluate() error {

	return nil
}