package db

// Options 参数
type Options struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Debug    bool   `json:"debug,omitempty"`
	LogLevel int    `json:"log_level"`
}

func Default() Options {
	return Options{
		Type:     "mysql",
		URL:      "root:root@tcp(git.zgwit.com:3306)/master?charset=utf8",
		Debug:    false,
		LogLevel: 4,
	}
}
