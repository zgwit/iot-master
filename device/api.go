package device

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/history"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/protocol"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {

	api.Register("POST", "/device/count", curd.ApiCount[Device]())

	//api.Register("POST","/search", curd.ApiSearch[Device]("id", "name", "product_id", "disabled", "created"))

	api.Register("POST", "/device/search", curd.ApiSearchWith[Device]([]*curd.With{
		{"project", "project_id", "id", "name", "project"},
		{"product", "product_id", "id", "name", "product"},
	}, "id", "name", "project_id", "product_id", "product_version", "disabled", "created"))

	api.Register("GET", "/device/list", curd.ApiList[Device]())

	api.Register("POST", "/device/create", curd.ApiCreateHook[Device](curd.GenerateID[Device](), nil))

	api.Register("GET", "/device/:id", curd.ParseParamStringId, curd.ApiGet[Device]())

	api.Register("POST", "/device/:id", curd.ParseParamStringId, curd.ApiUpdate[Device]())

	api.Register("GET", "/device/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Device](nil, nil))

	api.Register("GET", "/device/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[Device](true, nil, nil))

	api.Register("GET", "/device/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[Device](false, nil, nil))

	api.Register("GET", "/device/:id/values", curd.ParseParamStringId, deviceValues)

	api.Register("POST", "/device/:id/values", curd.ParseParamStringId, deviceValuesWrite)

	api.Register("GET", "/device/:id/history/:name", curd.ParseParamStringId, deviceHistory)

	api.Register("POST", "/device/:id/parameters", curd.ParseParamStringId, deviceParameters)

	api.Register("GET", "/device/statistic", deviceStatistic)

	api.Register("GET", "/device/:id/stations", curd.ParseParamStringId, deviceStations)

	api.Register("POST", "/device/:id/stations", curd.ParseParamStringId, deviceStationsUpdate)
}

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
// @Success 200 {object} curd.ReplyList[Device] 返回设备信息
// @Router /device/search [post]
func noopDeviceSearch() {}

// @Summary 查询设备
// @Schemes
// @Description 查询设备
// @Tags device
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Device] 返回设备信息
// @Router /device/list [get]
func noopDeviceList() {}

// @Summary 创建设备
// @Schemes
// @Description 创建设备
// @Tags device
// @Param search body Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Device] 返回设备信息
// @Router /device/create [post]
func noopDeviceCreate() {}

// @Summary 获取设备
// @Schemes
// @Description 获取设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Device] 返回设备信息
// @Router /device/{id} [get]
func noopDeviceGet() {}

// @Summary 修改设备
// @Schemes
// @Description 修改设备
// @Tags device
// @Param id path int true "设备ID"
// @Param device body Device true "设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Device] 返回设备信息
// @Router /device/{id} [post]
func noopDeviceUpdate() {}

// @Summary 删除设备
// @Schemes
// @Description 删除设备
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Device] 返回设备信息
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
// @Success 200 {object} curd.ReplyData[[]history.Point] 返回报警信息
// @Router /device/{id}/history/{name} [get]
func noopDeviceHistory() {}

// @Summary 修改设备参数
// @Schemes
// @Description 修改设备参数
// @Tags device
// @Param id path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Device] 返回设备信息
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
	obj.Total, err = db.Engine.Count(Device{})
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	obj.Online = GetOnlineCount()
	obj.Offline = obj.Total - obj.Online
	curd.OK(ctx, &obj)
}

func deviceValues(ctx *gin.Context) {
	dev := Get(ctx.GetString("id"))
	if dev == nil {
		curd.Fail(ctx, "设备未上线")
		return
	}
	curd.OK(ctx, dev.Values())
}

func deviceValuesWrite(ctx *gin.Context) {
	var values map[string]any
	err := ctx.ShouldBindJSON(&values)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	dev := Get(ctx.GetString("id"))
	if dev == nil {
		curd.Fail(ctx, "设备未上线")
		return
	}

	err = dev.WriteMany(values)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

func deviceHistory(ctx *gin.Context) {
	var dev Device
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
	var body map[string]any
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	dev := Device{Parameters: body}
	_, err = db.Engine.ID(ctx.GetString("id")).Cols("parameters").Update(&dev)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//TODO 重置设备
	//devices.Delete(ctx.GetString("id"))

	curd.OK(ctx, nil)
}

func deviceStations(ctx *gin.Context) {
	id := ctx.GetString("id")

	var dev Device
	_, err := db.Engine.ID(id).Get(&dev)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var prod product.Product
	_, err = db.Engine.ID(id).Get(&prod)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	p, err := protocol.Get(prod.Protocol)
	var cols []string
	for _, s := range p.Stations {
		cols = append(cols, s.Key)
	}

	var stations map[string]any
	_, err = db.Engine.Table(new(Device)).ID(id).Cols(cols...).Get(&stations)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, stations)
}

func deviceStationsUpdate(ctx *gin.Context) {
	id := ctx.GetString("id")

	var stations map[string]any
	err := ctx.BindJSON(&stations)

	var cols []string
	for k, _ := range stations {
		cols = append(cols, k)
	}

	_, err = db.Engine.Table(new(Device)).ID(id).Cols(cols...).Update(&stations)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}
