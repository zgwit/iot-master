package interval

import "github.com/zgwit/iot-master/internal/aggregator"

type ProjectDevice struct {
	Id   int    `json:"id" storm:"id,increment"`
	Name string `json:"name"`
}

type Project struct {
	Id       int  `json:"id"`
	Disabled bool `json:"disabled"`

	Devices []ProjectDevice `json:"devices"`

	Aggregators []aggregator.Aggregator `json:"aggregators"`
	Commands    []Command               `json:"commands"`
	Rectors     []Reactor               `json:"rectors"`
	Jobs        []Job                   `json:"jobs"`

	Context Context `json:"context"`

	deviceIndex map[string]interface{}
}

func (c *Project) Start() error {
	return nil
}

func (c *Project) Stop() error {
	return nil
}
