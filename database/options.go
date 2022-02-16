package database

//Options 参数
type Options struct {
	Path string `yaml:"path"`
}

func DefaultOptions() *Options {
	return &Options{
		Path: "data",
	}
}