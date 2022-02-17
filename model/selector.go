package model

//Selector 选择器
type Selector struct {
	IDs   []int    `json:"ids,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	Names []string `json:"names,omitempty"` //name
}
