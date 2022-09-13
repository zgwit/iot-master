package web

// Options 参数
type Options struct {
	Addr     string `yaml:"addr" json:"addr"`
	Debug    bool   `yaml:"debug" json:"debug"`
	Compress bool   `yaml:"compress" json:"compress"`
}
