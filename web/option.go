package web

//Options 参数
type Options struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug,omitempty"`
}

//DefaultOptions 默认
func DefaultOptions() *Options {
	return &Options{
		Addr:  ":8080",
		Debug: false,
	}
}
