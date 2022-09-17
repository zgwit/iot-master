package api

import (
	"encoding/json"
	"github.com/zgwit/iot-master/v2/internal/core"
	"github.com/zgwit/iot-master/v2/model"
)

func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)

	payload, err := json.Marshal(tunnel)
	if err != nil {
		return err
	}
	return core.Publish("/gateway/"+tunnel.GatewayId+"/download/tunnel", payload)
}

func afterTunnelUpdate(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	payload, err := json.Marshal(tunnel)
	if err != nil {
		return err
	}
	return core.Publish("/gateway/"+tunnel.GatewayId+"/download/tunnel", payload)
}

func afterTunnelDelete(id interface{}) error {
	gid := id.(string)
	core.TunnelStatus.Delete(id.(string))
	return core.Publish("/tunnel/"+gid+"/command/delete", []byte(""))
}
