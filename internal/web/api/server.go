package api

import (
	"encoding/json"
	"github.com/zgwit/iot-master/v3/internal/core"
	"github.com/zgwit/iot-master/v3/model"
)

func afterServerCreate(data interface{}) error {
	server := data.(*model.Server)

	payload, err := json.Marshal(server)
	if err != nil {
		return err
	}
	return core.Publish("/gateway/"+server.GatewayId+"/download/server", payload)
}

func afterServerUpdate(data interface{}) error {
	server := data.(*model.Server)
	payload, err := json.Marshal(server)
	if err != nil {
		return err
	}
	return core.Publish("/gateway/"+server.GatewayId+"/download/server", payload)
}

func afterServerDelete(id interface{}) error {
	gid := id.(string)
	core.ServerStatus.Delete(gid)
	return core.Publish("/server/"+gid+"/command/delete", []byte(""))
}
