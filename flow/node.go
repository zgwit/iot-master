package flow

type Node interface {
	Do() error
}

type baseNode struct {
	In  []string
	Out []string
}
