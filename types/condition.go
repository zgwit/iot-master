package types

type Evaluable interface {
	Evaluate() bool
}

type Condition struct {
	Both     bool        `json:"both"`
	Compares []Compare   `json:"compares"`
	Children []Condition `json:"children"`
}

func (cond *Condition) Evaluate() bool {
	if cond.Both {
		for _, it := range cond.Compares {
			if !it.Evaluate() {
				return false
			}
		}
		for _, it := range cond.Children {
			if !it.Evaluate() {
				return false
			}
		}
		return true
	} else {
		for _, it := range cond.Compares {
			if it.Evaluate() {
				return true
			}
		}
		for _, it := range cond.Children {
			if it.Evaluate() {
				return true
			}
		}
		return false
	}
}
