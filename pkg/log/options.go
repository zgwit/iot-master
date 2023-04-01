package log

type Output struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size,omitempty"`
	MaxAge     int    `json:"max_age,omitempty"`
	MaxBackups int    `json:"max_backups,omitempty"`
}

// Options 参数
type Options struct {
	Level  string `json:"level"`
	Caller bool   `json:"caller,omitempty"`
	Text   bool   `json:"text,omitempty"`
	Format string `json:"format,omitempty"`
	Output Output `json:"output"`
}

func Default() Options {
	return Options{
		Level:  "trace",
		Caller: true,
		Text:   true,
	}
}
