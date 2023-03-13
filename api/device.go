package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func deviceValues(ctx *gin.Context) {
	device := internal.Devices.Load(ctx.GetString("id"))
	if device == nil {
		curd.Fail(ctx, "设备未上线")
		return
	}
	curd.OK(ctx, device.Values)
}

func deviceParameters(ctx *gin.Context) {
	var body map[string]float64
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	dev := model.Device{Parameters: body}
	_, err = db.Engine.ID(ctx.GetString("id")).Update(&dev)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//TODO 重置设备
	internal.Devices.Delete(ctx.GetString("id"))

	curd.OK(ctx, nil)
}
