package interval

type Collector struct {
	Disabled bool
	Type     string //interval, clock, crontab
	Interval int
	Clock    int
	Crontab  string

	Code    int
	Address int
	Length  int

	//TODO Filters

}

func (c *Collector) Start() error {
	return nil
}


func (c *Collector) Stop() error {
	return nil
}