package config

//Web 参数
type Web struct {
	Addr     string `yaml:"addr"`
	Debug    bool   `yaml:"debug,omitempty"`
	Compress bool   `json:"compress"`
}

var WebDefault = Web{
	Addr:     ":8080",
	Debug:    false,
	Compress: true,
}
