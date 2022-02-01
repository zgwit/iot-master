package interval

type Directive struct {
	Value         float64 `json:"value"`
	ArgumentIndex int     `json:"argument_index"`
	Delay         int64   `json:"delay"`

	Device string `json:"device"`
	Point  string `json:"point"`

	device *Device
	point  *Point
}

func (d *Directive) Execute(argv []float64) error {

	return nil
}
