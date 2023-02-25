package internal

type Subset struct {
	Id         string
	Properties map[string]any
}

func NewSubset(id string) *Subset {
	return &Subset{
		Id:         id,
		Properties: make(map[string]any),
	}
}
