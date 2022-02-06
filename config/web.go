package config

type Web struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug,omitempty"`
}
