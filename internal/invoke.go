package interval

type Invoke struct {
	Command string `json:"command"`
	Argv []float64 `json:"argv"`

	command* Command
}

func (i Invoke) Execute() error {
	return i.command.Execute(i.Argv)
}