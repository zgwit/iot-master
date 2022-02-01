package interval

type ProjectDevice struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Project struct {
	Disabled bool `json:"disabled"`

	Devices []ProjectDevice `json:"devices"`

	Aggregators []Aggregator `json:"aggregators"`
	Commands    []Command    `json:"commands"`
	Rectors     []Reactor    `json:"rectors"`
	Jobs        []Job        `json:"jobs"`

	Context Context `json:"context"`

}

func (c *Project) Start() error {
	return nil
}

func (c *Project) Stop() error {
	return nil
}
