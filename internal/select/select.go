package _select


type Select struct {
	Names []string `json:"device,omitempty"` //name
	Ids   []int    `json:"ids,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

