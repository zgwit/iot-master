package config

// Web 参数
type Web struct {
	Addr  string `yaml:"addr" json:"addr"`
	Debug bool   `yaml:"debug,omitempty" json:"debug,omitempty"`
	Cors  bool   `json:"cors,omitempty" json:"cors,omitempty"`
	Gzip  bool   `json:"gzip,omitempty" json:"gzip,omitempty"`
}
