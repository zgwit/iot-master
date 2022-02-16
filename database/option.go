package database

//Option 参数
type Option struct {
	Path string `yaml:"path"`
}

func Default() *Option {
	return &Option{
		Path: "data",
	}
}