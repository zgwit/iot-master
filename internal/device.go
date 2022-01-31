package interval

type Device struct {
	Disabled bool

	Id   string
	Name string
	Tags []string

	Slave int

	//context
	Points      []Point
	Collectors  []Collector
	Calculators []Calculator
	Commands    []Command
	Alarms      []Alarm
	Jobs        []Job
}

func (c *Device) Start() error {
	return nil
}

func (c *Device) Stop() error {
	return nil
}
