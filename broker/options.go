package broker

type Options struct {
	Enable bool   `json:"enable"`
	Type   string `json:"type"`
	Addr   string `json:"addr"`
}
