package model

type Select struct {
	Ids   []int    `json:"ids,omitempty"`
	Names []string `json:"names,omitempty"` //name
	Tags  []string `json:"tags,omitempty"`
}
