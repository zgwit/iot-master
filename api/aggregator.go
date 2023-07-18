package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/model"
	"github.com/zgwit/iot-master/v4/pkg/curd"
)

// @Summary 查询聚合器数量
// @Schemes
// @Description 查询聚合器数量
// @Tags aggregator
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回聚合器数量
// @Router /aggregator/count [post]
func noopAggregatorCount() {}

// @Summary 查询聚合器
// @Schemes
// @Description 查询聚合器
// @Tags aggregator
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Aggregator] 返回聚合器信息
// @Router /aggregator/search [post]
func noopAggregatorSearch() {}

// @Summary 查询聚合器
// @Schemes
// @Description 查询聚合器
// @Tags aggregator
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Aggregator] 返回聚合器信息
// @Router /aggregator/list [get]
func noopAggregatorList() {}

// @Summary 创建聚合器
// @Schemes
// @Description 创建聚合器
// @Tags aggregator
// @Param search body model.Aggregator true "聚合器信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Aggregator] 返回聚合器信息
// @Router /aggregator/create [post]
func noopAggregatorCreate() {}

// @Summary 修改聚合器
// @Schemes
// @Description 修改聚合器
// @Tags aggregator
// @Param id path int true "聚合器ID"
// @Param aggregator body model.Aggregator true "聚合器信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Aggregator] 返回聚合器信息
// @Router /aggregator/{id} [post]
func noopAggregatorUpdate() {}

// @Summary 获取聚合器
// @Schemes
// @Description 获取聚合器
// @Tags aggregator
// @Param id path int true "聚合器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Aggregator] 返回聚合器信息
// @Router /aggregator/{id} [get]
func noopAggregatorGet() {}

// @Summary 删除聚合器
// @Schemes
// @Description 删除聚合器
// @Tags aggregator
// @Param id path int true "聚合器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Aggregator] 返回聚合器信息
// @Router /aggregator/{id}/delete [get]
func noopAggregatorDelete() {}

// @Summary 启用聚合器
// @Schemes
// @Description 启用聚合器
// @Tags aggregator
// @Param id path int true "聚合器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Aggregator] 返回聚合器信息
// @Router /aggregator/{id}/enable [get]
func noopAggregatorEnable() {}

// @Summary 禁用聚合器
// @Schemes
// @Description 禁用聚合器
// @Tags aggregator
// @Param id path int true "聚合器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Aggregator] 返回聚合器信息
// @Router /aggregator/{id}/disable [get]
func noopAggregatorDisable() {}

// @Summary 导出聚合器
// @Schemes
// @Description 导出聚合器
// @Tags aggregator
// @Accept json
// @Produce octet-stream
// @Router /aggregator/export [get]
func noopAggregatorExport() {}

// @Summary 导入聚合器
// @Schemes
// @Description 导入聚合器
// @Tags aggregator
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回聚合器数量
// @Router /aggregator/import [post]
func noopAggregatorImport() {}

func aggregatorRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Aggregator]())
	app.POST("/search", curd.ApiSearch[model.Aggregator]())
	app.GET("/list", curd.ApiList[model.Aggregator]())
	app.POST("/create", curd.ApiCreateHook[model.Aggregator](curd.GenerateRandomId[model.Aggregator](8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Aggregator]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.Aggregator](nil, nil,
		"id", "product_id", "name", "desc", "crontab", "aggregators", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.Aggregator](nil, nil))
	app.GET("/export", curd.ApiExport("aggregator", "aggregator"))
	app.POST("/import", curd.ApiImport("aggregator"))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[model.Aggregator](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[model.Aggregator](false, nil, nil))
}
