package db

// Options 参数
type Options struct {
	Type     string `yaml:"type" json:"type"`
	URL      string `yaml:"url" json:"url"`
	Debug    bool   `yaml:"debug,omitempty" json:"debug,omitempty"`
	LogLevel int    `json:"log_level" json:"log_level"`
}

func Default() Options {
	return Options{
		Type:     "mysql",
		URL:      "root:root@tcp(git.zgwit.com:3306)/master?charset=utf8",
		Debug:    false,
		LogLevel: 4,
	}
}
