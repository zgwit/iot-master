package api

import (
	"encoding/json"
	"errors"
	"github.com/zgwit/iot-master/internal/broker"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
)

func getServerGateway(TunnelId string) (string, error) {
	var tunnel model.Server
	has, err := db.Engine.Get(TunnelId, &tunnel)
	if err != nil {
		return "", err
	}
	if !has {
		return "", errors.New("找不到服务")
	}

	return tunnel.GatewayId, nil
}

func afterServerCreate(data interface{}) error {
	server := data.(*model.Server)

	gid, err := getTunnelGateway(server.Id)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(server)
	broker.MQTT.Publish("/gateway/"+gid+"/download/server", 0, false, payload)
	return err
}

func afterServerUpdate(data interface{}) error {
	server := data.(*model.Server)

	gid, err := getTunnelGateway(server.Id)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(server)
	broker.MQTT.Publish("/gateway/"+gid+"/download/server", 0, false, payload)
	return err
}

func afterServerDelete(id interface{}) error {
	core.ServerStatus.Delete(id.(string))
	return nil
}
