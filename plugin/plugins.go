package plugin

import (
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/log"
	"os"
	"path/filepath"
)

var internals []*Manifest

//func Register(m *Manifest)  {
//
//}

var plugins lib.Map[Plugin]

func New(manifest *Manifest) *Plugin {
	return &Plugin{
		Manifest: manifest,
		//values: map[string]float64{},
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
	//fn := fmt.Sprintf("%s/plugin/%s/manifest.yaml", viper.GetString("data"), id)
	fn := filepath.Join(viper.GetString("data"), "plugin", id, "manifest.yaml")

	var m Manifest
	err := lib.LoadYaml(fn, &m)
	if err != nil {
		return err
	}

	return From(id, &m)
}

func Store(id string, m *Manifest) error {
	//fn := fmt.Sprintf("%s/plugin/%s/manifest.yaml", viper.GetString("data"), id)
	fn := filepath.Join(viper.GetString("data"), "plugin", id, "manifest.yaml")
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

func Boot() error {
	//开机加载所有插件
	root := filepath.Join(viper.GetString("data"), "plugin")
	files, err := os.ReadDir(root)
	if err != nil {
		//return err
		log.Error(err)
		return nil
	}

	for _, stat := range files {
		info, _ := stat.Info()
		if info.IsDir() {
			err = Load(info.Name())
			if err != nil {
				log.Error(err)
			}
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

func GetPlugins() []*Manifest {
	var ps []*Manifest
	ps = append(ps, internals...)
	plugins.Range(func(id string, plugin *Plugin) bool {
		ps = append(ps, plugin.Manifest)
		return true
	})
	return ps
}
