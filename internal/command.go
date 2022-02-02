package interval

type Command struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Argc  int    `json:"argc,omitempty"`

	Directives []Directive `json:"directives"`
}

func (c Command) Execute(argv []float64) error {
	for _, d := range c.Directives {
		err := d.Execute(argv)
		if err != nil {
			return err
		}
	}
	return nil
}


type Directive struct {
	Value      float64 `json:"value,omitempty"`
	Arg        int     `json:"arg,omitempty"` //1 2 3
	Expression string  `json:"expression,omitempty"`
	expression *Expression

	Delay int64 `json:"delay"`

	//TODO 表达式

	Device string   `json:"device,omitempty"` //name
	Tags   []string `json:"tags,omitempty"`
	Point  string   `json:"point"`

	devices *[]Device
}

func (d *Directive) Execute(argv []float64) error {

	return nil
}


type Invoke struct {
	Command string    `json:"command"`
	Argv    []float64 `json:"argv"`

	command *Command
}

func (i Invoke) Execute() error {
	return i.command.Execute(i.Argv)
}
