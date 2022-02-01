package interval

type Command struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Argc  int    `json:"argc"`

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
