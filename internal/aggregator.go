package interval

import (
	"fmt"
	"strings"
)

type AggregatorType int

const (
	AggregatorNONE AggregatorType = iota
	AggregatorSUM
	AggregatorCOUNT
	AggregatorAVG
	AggregatorMEDIAN
	AggregatorMIN
	AggregatorMAX
	AggregatorFIRST
	AggregatorLAST
)

func (ct AggregatorType) Parse(tp string) error {
	strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "sum":
		ct = AggregatorSUM
	case "count":
		ct = AggregatorCOUNT
	case "avg":
		ct = AggregatorAVG
	case "median":
		ct = AggregatorMEDIAN
	case "min":
		ct = AggregatorMIN
	case "max":
		ct = AggregatorMAX
	case "first":
		ct = AggregatorFIRST
	case "last":
		ct = AggregatorLAST
	default:
		return fmt.Errorf("Unknown compare type: %s ", tp)
	}
	return nil
}

func (ct AggregatorType) String() string {
	var str string
	switch ct {
	case AggregatorSUM:
		str = "sum"
	case AggregatorCOUNT:
		str = "count"
	case AggregatorAVG:
		str = "avg"
	case AggregatorMEDIAN:
		str = "median"
	case AggregatorMIN:
		str = "min"
	case AggregatorMAX:
		str = "max"
	case AggregatorFIRST:
		str = "first"
	case AggregatorLAST:
		str = "last"
	default:
		str = "none"
	}
	return str
}

func (ct AggregatorType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.String() + `"`), nil
}

func (ct AggregatorType) UnmarshalJSON(buf []byte) error {
	return ct.Parse(string(buf))
}

type Aggregator struct {
	Type AggregatorType `json:"type"`
	As   string         `json:"as"`

	//TODO add device
	Tags []string `json:"tags"`
}
