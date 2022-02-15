package web

//Option 参数
type Option struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug,omitempty"`
}
