package model

//Aggregator 聚合器
type Aggregator struct {
	Type       string   `json:"type"`
	As         string   `json:"as"`
	From       string   `json:"from"`
	Select     Selector `json:"select"`
	Expression string   `json:"expression"`
}
