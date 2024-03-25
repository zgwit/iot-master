package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/web/curd"
)

// @Summary 查询网关数量
// @Schemes
// @Description 查询网关数量
// @Tags gateway
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回网关数量
// @Router /gateway/count [post]
func noopGatewayCount() {}

// @Summary 查询网关
// @Schemes
// @Description 查询网关
// @Tags gateway
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Gateway] 返回网关信息
// @Router /gateway/search [post]
func noopGatewaySearch() {}

// @Summary 查询网关
// @Schemes
// @Description 查询网关
// @Tags gateway
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Gateway] 返回网关信息
// @Router /gateway/list [get]
func noopGatewayList() {}

// @Summary 创建网关
// @Schemes
// @Description 创建网关
// @Tags gateway
// @Param search body types.Gateway true "网关信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Gateway] 返回网关信息
// @Router /gateway/create [post]
func noopGatewayCreate() {}

// @Summary 修改网关
// @Schemes
// @Description 修改网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Param gateway body types.Gateway true "网关信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Gateway] 返回网关信息
// @Router /gateway/{id} [post]
func noopGatewayUpdate() {}

// @Summary 获取网关
// @Schemes
// @Description 获取网关
// @Tags gateway
// @Param id path string true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Gateway] 返回网关信息
// @Router /gateway/{id} [get]
func noopGatewayGet() {}

// @Summary 删除网关
// @Schemes
// @Description 删除网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Gateway] 返回网关信息
// @Router /gateway/{id}/delete [get]
func noopGatewayDelete() {}

// @Summary 启用网关
// @Schemes
// @Description 启用网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Gateway] 返回网关信息
// @Router /gateway/{id}/enable [get]
func noopGatewayEnable() {}

// @Summary 禁用网关
// @Schemes
// @Description 禁用网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Gateway] 返回网关信息
// @Router /gateway/{id}/disable [get]
func noopGatewayDisable() {}

func gatewayRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[broker.Gateway]())
	//app.POST("/search", curd.ApiSearch[types.Gateway]())
	app.POST("/search", curd.ApiSearchWith[broker.Gateway]([]*curd.With{
		{"project", "project_id", "id", "name", "project"},
	}, "id", "name", "project_id", "disabled", "created"))
	app.GET("/list", curd.ApiList[broker.Gateway]())
	app.POST("/create", curd.ApiCreateHook[broker.Gateway](curd.GenerateID[broker.Gateway](), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[broker.Gateway]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[broker.Gateway](nil, nil,
		"id", "name", "description", "project_id", "username", "password", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[broker.Gateway](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[broker.Gateway](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[broker.Gateway](false, nil, nil))
}
