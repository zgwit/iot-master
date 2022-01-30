package types

type Command struct {
	Name  string
	Label string
	Argc  int

	Directives []Directive
}

func (c Command) Execute(argv []float64) error {

}

