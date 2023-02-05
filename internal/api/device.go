package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
)

func afterDeviceCreate(data interface{}) error {
	device := data.(*model.Device)

	internal.Devices.Store(device.Id, &internal.Device{
		Id:         device.Id,
		Properties: make(map[string]any),
	})

	payload, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return internal.Publish("/gateway/"+device.GatewayId+"/download/device", payload)
}

func afterDeviceUpdate(data interface{}) error {
	device := data.(*model.Device)

	payload, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return internal.Publish("/gateway/"+device.GatewayId+"/download/device", payload)
}

func afterDeviceDelete(id interface{}) error {
	did := id.(string)
	internal.Devices.Delete(did)
	return internal.Publish("/device/"+did+"/command/delete", []byte(""))
}

func deviceProperties(ctx *gin.Context) {
	device := internal.Devices.Load(ctx.GetString("id"))
	if device == nil {
		replyFail(ctx, "找不到设备变量")
		return
	}
	replyOk(ctx, device.Properties)
}
