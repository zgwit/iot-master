package types

type Invoke struct {
	Command string
	Argv []float64

	command* Command
}

func (i Invoke) Execute() error {
	return i.command.Execute(i.Argv)
}