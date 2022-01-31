package interval

type Directive struct {
	Value         float64
	ArgumentIndex int
	Delay         int64

	Device string
	Point  string

	device *Device
	point  *Point
}

func (d *Directive) Execute(argv []float64) error {



	return nil
}