package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var ROOT string

func init() {
	var err error
	ROOT, err = os.UserConfigDir()
	if err != nil {
		ROOT = ""
	} else {
		ROOT = filepath.Join(ROOT, "iot-master")
	}
}

func Load(name string, cfg any) error {
	fn := filepath.Join(ROOT, name)
	y, err := os.Open(fn)
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

func Store(name string, cfg any) error {
	fn := filepath.Join(ROOT, name)
	y, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755) //os.Create(name)
	if err != nil {
		return err
	}
	defer y.Close()

	e := yaml.NewEncoder(y)
	defer e.Close()

	err = e.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
