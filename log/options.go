package log

type Output struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

//Options 参数
type Options struct {
	Development bool   `yaml:"development"`
	Format      string `yaml:"format,omitempty"`
	Level       string `yaml:"level"`
	Output      Output `yaml:"output"`
}

func DefaultOptions() *Options {
	return &Options{
		Development: false,
		Level:       "debug",
		Output: Output{
			Filename:   "log.txt",
		},
	}
}
