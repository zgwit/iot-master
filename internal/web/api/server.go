package api

import (
	"github.com/zgwit/iot-master/model"
)

func afterServerCreate(data interface{}) error {
	server := data.(*model.Server)
	if !server.Disabled {
		return core.LoadServer(server.Id)
	}
	return nil
}

func afterServerUpdate(data interface{}) error {
	server := data.(*model.Server)
	_ = core.RemoveServer(server.Id)
	return core.LoadServer(server.Id)
}

func afterServerDelete(id interface{}) error {
	return core.RemoveServer(id.(int64))
}

func afterServerEnable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}

func afterServerDisable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}
