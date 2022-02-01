package interval

type Directive struct {
	Value      float64 `json:"value,omitempty"`
	Arg        int     `json:"arg,omitempty"` //1 2 3
	Expression string  `json:"expression,omitempty"`
	expression *Expression

	Delay int64 `json:"delay"`

	//TODO 表达式

	Device string   `json:"device"` //name
	Tags   []string `json:"tags"`
	Point  string   `json:"point"`

	devices *[]Device
}

func (d *Directive) Execute(argv []float64) error {

	return nil
}
