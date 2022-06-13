package config

type LogOutput struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

//Log 参数
type Log struct {
	Text   bool      `yaml:"text"`
	Format string    `yaml:"format,omitempty"`
	Level  string    `yaml:"level"`
	Output LogOutput `yaml:"output"`
}

var LogDefault = Log{
	Text:  false,
	Level: "error",
	Output: LogOutput{
		Filename: "log.txt",
	},
}
