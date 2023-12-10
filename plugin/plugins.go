package plugin

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/app"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
)

var plugins lib.Map[Plugin]

func New(model *app.Model) *Plugin {
	return &Plugin{
		Model: model,
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
	var p app.Model
	get, err := db.Engine.ID(id).Get(&p)
	if err != nil {
		return err
	}
	if !get {
		return fmt.Errorf("plugin %s not found", id)
	}

	return From(&p)
}

func From(model *app.Model) error {
	p := New(model)

	err := p.Start()
	if err != nil {
		return err
	}

	plugins.Store(model.Id, p)

	return nil
}

func LoadAll() error {
	//开机加载所有插件
	var ps []*app.Model
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		if p.Disabled {
			continue
		}
		err = From(p)
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
