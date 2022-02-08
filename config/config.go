package config

type Configure struct {
	Web      Web      `yaml:"web"`
	Database Database `yaml:"database"`
	History  History  `yaml:"history"`
}

var Config Configure = Configure{
	Web: Web{
		Addr: ":8080",
	},
	Database: Database{
		Path: ".",
	},
	History: History{
		DataPath: ".",
	},
}

func Load() error {

	return nil
}
