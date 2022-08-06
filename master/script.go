package master

import (
	"github.com/dop251/goja"
	"iot-master/model"
)

//Script 脚本
type Script struct {
	model.Script

	vm     *goja.Runtime
	onData goja.Callable
}

func (s *Script) Init(ctx map[string]interface{}) error {
	s.vm = goja.New()
	_, err := s.vm.RunString(s.Source)
	if err != nil {
		return err
	}
	err = s.vm.Set("context", ctx)
	if err != nil {
		return err
	}
	s.onData, _ = goja.AssertFunction(s.vm.Get("onData"))
	return err
}

func (s *Script) OnData(ctx map[string]interface{}) error {
	_, err := s.onData(goja.Undefined(), s.vm.ToValue(ctx))
	return err
}
