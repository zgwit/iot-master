package config

type LogOutput struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

//Log 参数
type Log struct {
	Debug  bool      `yaml:"debug"`
	Format string    `yaml:"format,omitempty"`
	Level  string    `yaml:"level"`
	Output LogOutput `yaml:"output"`
}

var LogDefault = Log{
	Debug: false,
	Level: "debug",
	Output: LogOutput{
		Filename: "log.txt",
	},
}
