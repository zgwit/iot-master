package interval

type Command struct {
	Name  string `json:"name"`
	Label string `json:"label,omitempty"`
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
	Point string  `json:"point"`
	Value float64 `json:"value,omitempty"`
	Delay int64   `json:"delay"`

	//使用参数
	Arg int `json:"arg,omitempty"` //0:默认参数 1 2 3

	//TODO 使用表达式
	Expression string `json:"expression,omitempty"`

	//目标设备（Project中使用）（不合适！！！）
	Device string   `json:"device,omitempty"` //name
	Tags   []string `json:"tags,omitempty"`

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
