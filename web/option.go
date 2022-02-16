package web

//Option 参数
type Option struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug,omitempty"`
}

//Default 默认
func Default() *Option {
	return &Option{
		Addr:  ":8080",
		Debug: false,
	}
}
