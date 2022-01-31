package interval

type Project struct {
	Device []string

	Aggregators []Aggregator
	Commands    []Command
	Alarms      []Alarm
	Jobs        []Job
	Rectors     []Rector
}


func (c *Project) Start() error {
	return nil
}

func (c *Project) Stop() error {
	return nil
}
