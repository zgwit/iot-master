package api

import (
	"encoding/json"
	"github.com/zgwit/iot-master/internal/broker"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/model"
)

func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)

	payload, err := json.Marshal(tunnel)
	broker.MQTT.Publish("/gateway/"+tunnel.GatewayId+"/download/tunnel", 0, false, payload)
	return err
}

func afterTunnelDelete(id interface{}) error {
	core.TunnelStatus.Delete(id.(string))
	return nil
}
