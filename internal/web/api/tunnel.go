package api

import (
	"encoding/json"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/model"
)

func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)

	payload, err := json.Marshal(tunnel)
	if err != nil {
		return err
	}
	return core.Publish("/gateway/"+tunnel.GatewayId+"/download/tunnel", payload)
}

func afterTunnelDelete(id interface{}) error {
	core.TunnelStatus.Delete(id.(string))
	return nil
}
