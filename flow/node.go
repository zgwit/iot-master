package flow

type Node interface {
	Do() error
}

type baseNode struct {
}
