package model

//Aggregator 聚合器
type Aggregator struct {
	Targets    []string `json:"targets"`
	Type       string   `json:"type"`
	As         string   `json:"as"`
	Expression string   `json:"expression"`
}
