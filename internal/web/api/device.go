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
	//重新启动
	_ = core.RemoveDevice(device.Id)
	dev, err := core.LoadDevice(device.Id)
	if err == nil {
		err = dev.Start()
	}
	return err
}

func afterDeviceDelete(id interface{}) error {
	return core.RemoveDevice(id.(int64))
}

func afterDeviceEnable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}

func afterDeviceDisable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}

func deviceContext(ctx *gin.Context) {
	device := core.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	replyOk(ctx, device.Context)
}

func deviceContextUpdate(ctx *gin.Context) {
	var values map[string]interface{}
	err := ctx.ShouldBindJSON(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	device := core.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}

	for k, v := range values {
		err := device.Set(k, v)
		if err != nil {
			replyError(ctx, err)
			return
		}
	}

	replyOk(ctx, nil)
}

func deviceRefresh(ctx *gin.Context) {
	device := core.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err := device.Refresh()
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, device.Context)
}
