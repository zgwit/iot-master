package model

//Aggregator 聚合器
type Aggregator struct {
	Targets    []string `json:"targets"`
	Type       string   `json:"type"`
	As         string   `json:"as"`
	From       string   `json:"from"`
	Expression string   `json:"expression"`
}
