package model

type Mapping struct {
	Points []*Point `json:"points"`
	Slave  int      `json:"slave"`
}
