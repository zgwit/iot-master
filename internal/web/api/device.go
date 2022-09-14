package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/broker"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
)

func getTunnelGateway(TunnelId string) (string, error) {
	var tunnel model.Tunnel
	has, err := db.Engine.Get(TunnelId, &tunnel)
	if err != nil {
		return "", err
	}
	if !has {
		return "", errors.New("找不到通道")
	}

	return tunnel.GatewayId, nil
}
func afterDeviceCreate(data interface{}) error {
	device := data.(*model.Device)

	core.Devices.Store(device.Id, core.NewDevice(device.Id))

	gid, err := getTunnelGateway(device.TunnelId)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(device)
	broker.MQTT.Publish("/gateway/"+gid+"/download/device", 0, false, payload)
	return err
}

func afterDeviceUpdate(data interface{}) error {
	device := data.(*model.Device)

	gid, err := getTunnelGateway(device.TunnelId)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(device)
	broker.MQTT.Publish("/gateway/"+gid+"/download/device", 0, false, payload)

	return err
}

func afterDeviceDelete(id interface{}) error {
	core.Devices.Delete(id.(string))
	return nil
}

func deviceValues(ctx *gin.Context) {
	device := core.Devices.Load(ctx.GetString("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	replyOk(ctx, device.Values)
}

func deviceAssign(ctx *gin.Context) {
	var values map[string]interface{}
	err := ctx.ShouldBindJSON(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	device := core.Devices.Load(ctx.GetString("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}

	err = device.Assign(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceRefresh(ctx *gin.Context) {
	device := core.Devices.Load(ctx.GetString("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err := device.Refresh()
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}
