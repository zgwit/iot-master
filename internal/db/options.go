package db

// Options 参数
type Options struct {
	Type     string `yaml:"type" json:"type"`
	URL      string `yaml:"url" json:"url"`
	Debug    bool   `yaml:"debug" json:"debug"`
	LogLevel int    `json:"log_level" json:"log_level"`
	Sync     bool   `yaml:"sync" json:"sync"`
}
