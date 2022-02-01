package interval

type Device struct {
	Disabled bool `json:"disabled"`

	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`

	//从机号
	Slave int `json:"slave"`

	Points      []Point      `json:"points"`
	Collectors  []Collector  `json:"collectors"`
	Calculators []Calculator `json:"calculators"`
	Commands    []Command    `json:"commands"`
	Reactors    []Reactor    `json:"reactors"`
	Jobs        []Job        `json:"jobs"`

	//上下文
	Context Context `json:"context"`

	//命令索引
	commandIndex map[string]*Command
}

func (c *Device) Start() error {
	return nil
}

func (c *Device) Stop() error {
	return nil
}
