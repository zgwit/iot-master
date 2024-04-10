package settings

import (
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/types"
)

type Module struct {
	Name   string           `json:"name"`
	Module string           `json:"module"`
	Title  string           `json:"title,omitempty"`
	Form   []types.FormItem `json:"-"`
}

var modules lib.Map[Module]

func Register(module string, form *Module) {
	modules.Store(module, form)
}

func Unregister(module string) {
	modules.Delete(module)
}
