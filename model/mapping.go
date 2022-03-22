package model

type Mapping struct {
	Points  []*Point `json:"points"`
	Station int      `json:"station"`
}
