package interval

type AggregatorType int

const (
	AggregatorTypeNone AggregatorType = iota
	AggregatorTypeSum
	AggregatorTypeCount
	AggregatorTypeAvg
	
)

type Aggregator struct {
	
}