package log

type Output struct {
	Filename   string `yaml:"filename" json:"filename"`
	MaxSize    int    `yaml:"max_size" json:"max_size,omitempty"`
	MaxAge     int    `yaml:"max_age" json:"max_age,omitempty"`
	MaxBackups int    `yaml:"max_backups" json:"max_backups,omitempty"`
}

// Options 参数
type Options struct {
	Level  string `yaml:"level" json:"level"`
	Caller bool   `yaml:"caller" json:"caller"`
	Text   bool   `yaml:"text" json:"text,omitempty"`
	Format string `yaml:"format,omitempty" json:"format,omitempty"`
	Output Output `yaml:"output" json:"output"`
}

func Default() Options {
	return Options{
		Level:  "trace",
		Caller: true,
		Text:   true,
	}
}
