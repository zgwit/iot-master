package web

type Option struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug,omitempty"`
}
