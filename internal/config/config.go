package config

type Configure struct {
	Web      Web      `yaml:"web"`
	Database Database `yaml:"database"`
	History  History  `yaml:"history"`
}

var Config Configure

func Load() error {

	return nil
}