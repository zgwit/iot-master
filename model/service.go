package model

type Service struct {
	Name string `json:"name"`
	Net  string `json:"net,omitempty"`
	Addr string `json:"addr"`
}
