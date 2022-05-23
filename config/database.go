package config

//Database 参数
type Database struct {
	Path string `yaml:"path"`
}

var DatabaseDefault = Database{
	Path: "data",
}
