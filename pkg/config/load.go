package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func Load(filename string, cfg any) error {
	y, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer y.Close()

	d := yaml.NewDecoder(y)
	err = d.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}
