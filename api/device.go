package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func deviceRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Device]())
	app.GET("/list", curd.ApiList[model.Device]())
	app.POST("/create", curd.ApiCreate[model.Device](curd.GenerateRandomKey(12), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Device]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Device](nil, nil,
		"id", "gateway_id", "product_id", "group_id", "type", "name", "desc", "username", "password", "parameters", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Device](nil, nil))

	app.GET("/:id/values", curd.ParseParamStringId, deviceValues)
	app.POST("/:id/parameters", curd.ParseParamStringId, deviceParameters)
}

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
