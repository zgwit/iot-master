package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/model"
)

func afterDeviceCreate(data interface{}) error {
	device := data.(*model.Device)
	//启动
	dev, err := core.LoadDevice(device.Id)
	if err == nil {
		err = dev.Start()
	}
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
	return core.RemoveDevice(id.(uint64))
}

func deviceStart(ctx *gin.Context) {
	device := core.GetDevice(ctx.GetUint64("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Start()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	device := core.GetDevice(ctx.GetUint64("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Stop()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterDeviceEnable(id interface{}) error {
	_ = core.RemoveDevice(id.(uint64))
	dev, err := core.LoadDevice(id.(uint64))
	if err != nil {
		err = dev.Start()
	}
	return err
}

func afterDeviceDisable(id interface{}) error {
	return core.RemoveDevice(id.(uint64))
}

func deviceContext(ctx *gin.Context) {
	device := core.GetDevice(ctx.GetUint64("id"))
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

	device := core.GetDevice(ctx.GetUint64("id"))
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
	device := core.GetDevice(ctx.GetUint64("id"))
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

func deviceRefreshPoint(ctx *gin.Context) {
	device := core.GetDevice(ctx.GetUint64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	val, err := device.RefreshPoint(ctx.Param("name"))
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, val)
}

type executeBody struct {
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments"`
}

func deviceExecute(ctx *gin.Context) {
	var body executeBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	device := core.GetDevice(ctx.GetUint64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err = device.Execute(body.Command, body.Arguments)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}
