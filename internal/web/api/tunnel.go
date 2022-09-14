package api

import (
	"github.com/zgwit/iot-master/model"
)

func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	if !tunnel.Disabled {
		return core.LoadTunnel(tunnel.Id)
	}
	return nil
}

func afterTunnelDelete(id interface{}) error {
	return core.RemoveTunnel(id.(int64))
}

func afterTunnelEnable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}

func afterTunnelDisable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}
