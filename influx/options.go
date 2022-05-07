package influx

//Options 参数
type Options struct {
	Enable      bool   `yaml:"enable"`
	Bucket      string `yaml:"bucket"`
	ORG         string `yaml:"org"`
	Token       string `yaml:"token"`
	URL         string `yaml:"url"`
	Measurement string `yaml:"measurement"`
}

func DefaultOptions() *Options {
	return &Options{}
}
