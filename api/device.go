package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/device"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

// @Summary 查询设备数量
// @Schemes
// @Description 查询设备数量
// @Tags device
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回设备数量
// @Router /device/count [post]
func noopDeviceCount() {}

// @Summary 查询设备
// @Schemes
// @Description 查询设备
// @Tags device
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Device] 返回设备信息
// @Router /device/search [post]
func noopDeviceSearch() {}

// @Summary 查询设备
// @Schemes
// @Description 查询设备
// @Tags device
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Device] 返回设备信息
// @Router /device/list [get]
func noopDeviceList() {}

// @Summary 创建设备
// @Schemes
// @Description 创建设备
// @Tags device
// @Param search body model.Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Device] 返回设备信息
// @Router /device/create [post]
func noopDeviceCreate() {}

// @Summary 获取设备
// @Schemes
// @Description 获取设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Device] 返回设备信息
// @Router /device/{id} [get]
func noopDeviceGet() {}

// @Summary 修改设备
// @Schemes
// @Description 修改设备
// @Tags device
// @Param id path int true "设备ID"
// @Param device body model.Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Device] 返回设备信息
// @Router /device/{id} [post]
func noopDeviceUpdate() {}

// @Summary 删除设备
// @Schemes
// @Description 删除设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Device] 返回设备信息
// @Router /device/{id}/delete [get]
func noopDeviceDelete() {}

// @Summary 导出设备
// @Schemes
// @Description 导出设备
// @Tags device
// @Accept json
// @Produce octet-stream
// @Router /device/export [get]
func noopDeviceExport() {}

// @Summary 导入设备
// @Schemes
// @Description 导入设备
// @Tags device
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回设备数量
// @Router /device/import [post]
func noopDeviceImport() {}

// @Summary 设备变量
// @Schemes
// @Description 设备变量
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Variables] 返回设备信息
// @Router /device/{id}/values [get]
func noopDeviceValues() {}

// @Summary 修改设备参数
// @Schemes
// @Description 修改设备参数
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Device] 返回设备信息
// @Router /device/{id}/parameters [post]
func noopDeviceParameters() {}

type deviceStatisticObj struct {
	Online  int64 `json:"online"`
	Offline int64 `json:"offline"`
	Total   int64 `json:"total"`
}

// @Summary 设备统计
// @Schemes
// @Description 设备统计
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[deviceStatisticObj] 返回设备信息
// @Router /device/statistic [get]
func deviceStatistic(ctx *gin.Context) {
	var obj deviceStatisticObj
	var err error
	obj.Total, err = db.Engine.Count(model.Device{})
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	obj.Online = device.GetOnlineCount()
	obj.Offline = obj.Total - obj.Online
	curd.OK(ctx, &obj)
}

func deviceRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Device]())

	app.POST("/search", curd.ApiSearchWith[model.Device]([]*curd.Join{
		{"product", "product_id", "id", "name", "product"},
	}, "id", "name", "product_id", "disabled", "created"))

	app.GET("/list", curd.ApiList[model.Device]())

	app.POST("/create", curd.ApiCreateHook[model.Device](curd.GenerateRandomId[model.Device](12), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Device]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.Device](nil, nil,
		"id", "gateway_id", "product_id", "group_id", "type", "name", "desc", "username", "password", "parameters", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.Device](nil, nil))

	app.GET("/export", curd.ApiExport("device", "设备"))

	app.POST("/import", curd.ApiImport("device"))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[model.Device](true, nil, nil))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[model.Device](false, nil, nil))

	app.GET("/:id/values", curd.ParseParamStringId, deviceValues)

	app.POST("/:id/parameters", curd.ParseParamStringId, deviceParameters)

	app.GET("/statistic", deviceStatistic)
}

func deviceValues(ctx *gin.Context) {
	dev := device.Get(ctx.GetString("id"))
	if dev == nil {
		curd.Fail(ctx, "设备未上线")
		return
	}
	curd.OK(ctx, dev.Values)
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
	//device.devices.Delete(ctx.GetString("id"))

	curd.OK(ctx, nil)
}
