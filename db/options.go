package db

// Options 参数
type Options struct {
	Type     string `yaml:"type" json:"type"`
	URL      string `yaml:"url" json:"url"`
	Debug    bool   `yaml:"debug,omitempty" json:"debug,omitempty"`
	Sync     bool   `yaml:"sync,omitempty" json:"sync,omitempty"`
	LogLevel int    `json:"log_level" json:"log_level"`
}
