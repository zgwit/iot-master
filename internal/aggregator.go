package interval

type AggregatorType int

const (
	AggregatorTypeNone AggregatorType = iota
	AggregatorTypeSum
	AggregatorTypeCount
	AggregatorTypeAverage
	AggregatorTypeMedian
	AggregatorTypeMin
	AggregatorTypeMax
	AggregatorTypeFirst
	AggregatorTypeLast
)



type Aggregator struct {
	
}