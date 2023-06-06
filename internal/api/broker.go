package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询总线数量
// @Schemes
// @Description 查询总线数量
// @Tags broker
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回总线数量
// @Router /broker/count [post]
func noopBrokerCount() {}

// @Summary 查询总线
// @Schemes
// @Description 查询总线
// @Tags broker
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Broker] 返回总线信息
// @Router /broker/search [post]
func noopBrokerSearch() {}

// @Summary 查询总线
// @Schemes
// @Description 查询总线
// @Tags broker
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Broker] 返回总线信息
// @Router /broker/list [get]
func noopBrokerList() {}

// @Summary 创建总线
// @Schemes
// @Description 创建总线
// @Tags broker
// @Param search body model.Broker true "总线信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Broker] 返回总线信息
// @Router /broker/create [post]
func noopBrokerCreate() {}

// @Summary 修改总线
// @Schemes
// @Description 修改总线
// @Tags broker
// @Param id path string true "总线ID"
// @Param broker body model.Broker true "总线信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Broker] 返回总线信息
// @Router /broker/{id} [post]
func noopBrokerUpdate() {}

// @Summary 获取总线
// @Schemes
// @Description 获取总线
// @Tags broker
// @Param id path string true "总线ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Broker] 返回总线信息
// @Router /broker/{id} [get]
func noopBrokerGet() {}

// @Summary 删除总线
// @Schemes
// @Description 删除总线
// @Tags broker
// @Param id path string true "总线ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Broker] 返回总线信息
// @Router /broker/{id}/delete [get]
func noopBrokerDelete() {}

// @Summary 启用总线
// @Schemes
// @Description 启用总线
// @Tags broker
// @Param id path string true "总线ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Broker] 返回总线信息
// @Router /broker/{id}/enable [get]
func noopBrokerEnable() {}

// @Summary 禁用总线
// @Schemes
// @Description 禁用总线
// @Tags broker
// @Param id path string true "总线ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Broker] 返回总线信息
// @Router /broker/{id}/disable [get]
func noopBrokerDisable() {}

// @Summary 导出总线
// @Schemes
// @Description 导出总线
// @Tags broker
// @Accept json
// @Produce octet-stream
// @Success 200 {object} curd.ReplyList[model.Broker] 返回压缩包
// @Router /broker/export [get]
func noopBrokerExport() {}

// @Summary 导入总线
// @Schemes
// @Description 导入总线
// @Tags broker
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回总线数量
// @Router /broker/import [post]
func noopBrokerImport() {}

func brokerRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Broker]())
	app.POST("/search", curd.ApiSearch[model.Broker]())
	app.GET("/list", curd.ApiList[model.Broker]())
	app.POST("/create", curd.ApiCreateHook[model.Broker](curd.GenerateRandomId[model.Broker](12), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Broker]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.Broker](nil, nil,
		"id", "name", "type", "port", "desc", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.Broker](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[model.Broker](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[model.Broker](false, nil, nil))
	app.GET("/export", curd.ApiExport("broker", "总线"))
	app.POST("/import", curd.ApiImport("broker"))
}

func afterBrokerCreate(data interface{}) error {
	//broker := data.(*model.Broker)

	//TODO start broker
	return nil
}

func afterBrokerUpdate(data interface{}) error {
	//broker := data.(*model.Broker)

	//TODO restart broker
	return nil
}

func afterBrokerDelete(id interface{}) error {
	//gid := id.(string)

	//todo stop broker
	return nil
}
