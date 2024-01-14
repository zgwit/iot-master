package plugin

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/types"
)

var plugins lib.Map[Plugin]

func New(manifest *Manifest) *Plugin {
	return &Plugin{
		Manifest: manifest,
		//Values: map[string]float64{},
	}
}

func Ensure(id string) (*Plugin, error) {
	p := plugins.Load(id)
	if p == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		p = plugins.Load(id)
	}
	return p, nil
}

func Get(id string) *Plugin {
	return plugins.Load(id)
}

func Load(id string) error {
	fn := fmt.Sprintf("plugin/%s/manifest.yaml", id)

	var m Manifest
	err := lib.LoadYaml(fn, &m)
	if err != nil {
		return err
	}

	return From(id, &m)
}

func Store(id string, m *Manifest) error {
	fn := fmt.Sprintf("plugin/%s/manifest.yaml", id)
	err := lib.StoreYaml(fn, m)
	if err != nil {
		return err
	}
	return From(id, m)
}

func From(id string, manifest *Manifest) error {
	p := New(manifest)

	err := p.Start()
	if err != nil {
		return err
	}

	plugins.Store(id, p)

	return nil
}

func LoadAll() error {
	//开机加载所有插件
	var ps []*types.Plugin
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		if p.Disabled {
			continue
		}
		err = Load(p.Id)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}

func Close() {
	plugins.Range(func(id string, plugin *Plugin) bool {
		_ = plugin.Close()
		return true
	})
}
