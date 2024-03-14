package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/history"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/curd"
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
// @Success 200 {object} curd.ReplyList[types.Device] 返回设备信息
// @Router /device/search [post]
func noopDeviceSearch() {}

// @Summary 查询设备
// @Schemes
// @Description 查询设备
// @Tags device
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Device] 返回设备信息
// @Router /device/list [get]
func noopDeviceList() {}

// @Summary 创建设备
// @Schemes
// @Description 创建设备
// @Tags device
// @Param search body types.Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Device] 返回设备信息
// @Router /device/create [post]
func noopDeviceCreate() {}

// @Summary 获取设备
// @Schemes
// @Description 获取设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Device] 返回设备信息
// @Router /device/{id} [get]
func noopDeviceGet() {}

// @Summary 修改设备
// @Schemes
// @Description 修改设备
// @Tags device
// @Param id path int true "设备ID"
// @Param device body types.Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Device] 返回设备信息
// @Router /device/{id} [post]
func noopDeviceUpdate() {}

// @Summary 删除设备
// @Schemes
// @Description 删除设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Device] 返回设备信息
// @Router /device/{id}/delete [get]
func noopDeviceDelete() {}

// @Summary 设备变量
// @Schemes
// @Description 设备变量
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Variables] 返回设备信息
// @Router /device/{id}/values [get]
func noopDeviceValues() {}

// @Summary 查询历史数据
// @Schemes
// @Description 查询历史数据
// @Tags device
// @Param id path string true "设备ID"
// @Param name path string true "变量名称"
// @Param start query string false "起始时间"
// @Param end query string false "结束时间"
// @Param window query string false "窗口时间"
// @Param method query string false "算法"
// @Produce json
// @Success 200 {object} ReplyData[[]history.Point] 返回报警信息
// @Router /device/{id}/history/{name} [get]
func noopDeviceHistory() {}

// @Summary 修改设备参数
// @Schemes
// @Description 修改设备参数
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Device] 返回设备信息
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
	obj.Total, err = db.Engine.Count(types.Device{})
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	obj.Online = device.GetOnlineCount()
	obj.Offline = obj.Total - obj.Online
	curd.OK(ctx, &obj)
}

func deviceRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Device]())

	//app.POST("/search", curd.ApiSearch[types.Device]("id", "name", "product_id", "disabled", "created"))

	app.POST("/search", curd.ApiSearchWith[types.Device]([]*curd.Join{
		{"project", "project_id", "id", "name", "project"},
		{"product", "product_id", "id", "name", "product"},
		{"gateway", "gateway_id", "id", "name", "gateway"},
	}, "id", "name", "project_id", "product_id", "product_version", "gateway_id", "disabled", "created"))

	app.GET("/list", curd.ApiList[types.Device]())

	app.POST("/create", curd.ApiCreateHook[types.Device](curd.GenerateID[types.Device](), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Device]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Device](nil, nil,
		"id", "gateway_id", "product_id", "product_version", "project_id", "type", "name", "description", "parameters", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Device](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Device](true, nil, nil))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Device](false, nil, nil))

	app.GET("/:id/values", curd.ParseParamStringId, deviceValues)

	app.GET("/:id/history/:name", curd.ParseParamStringId, deviceHistory)

	app.POST("/:id/parameters", curd.ParseParamStringId, deviceParameters)

	app.GET("/statistic", deviceStatistic)
}

func deviceValues(ctx *gin.Context) {
	dev := device.Get(ctx.GetString("id"))
	if dev == nil {
		curd.Fail(ctx, "设备未上线")
		return
	}
	curd.OK(ctx, dev.Values())
}

func deviceHistory(ctx *gin.Context) {
	var dev types.Device
	has, err := db.Engine.ID(ctx.GetString("id")).Get(&dev)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if !has {
		curd.Fail(ctx, "找不到设备")
		return
	}

	key := ctx.Param("name")
	start := ctx.DefaultQuery("start", "-5h")
	end := ctx.DefaultQuery("end", "0h")
	window := ctx.DefaultQuery("window", "10m")
	method := ctx.DefaultQuery("method", "mean") //last

	points, err := history.Query(dev.ProductId, dev.Id, key, start, end, window, method)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, points)
}

func deviceParameters(ctx *gin.Context) {
	var body map[string]float64
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	dev := types.Device{Parameters: body}
	_, err = db.Engine.ID(ctx.GetString("id")).Update(&dev)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//TODO 重置设备
	//device.devices.Delete(ctx.GetString("id"))

	curd.OK(ctx, nil)
}
