package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/curd"
	"github.com/zgwit/iot-master/v4/web/export"
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

// @Summary 导出网关
// @Schemes
// @Description 导出网关
// @Tags gateway
// @Accept json
// @Produce octet-stream
// @Success 200 {object} curd.ReplyList[types.Gateway] 返回压缩包
// @Router /gateway/export [get]
func noopGatewayExport() {}

// @Summary 导入网关
// @Schemes
// @Description 导入网关
// @Tags gateway
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回网关数量
// @Router /gateway/import [post]
func noopGatewayImport() {}

func gatewayRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Gateway]())
	//app.POST("/search", curd.ApiSearch[types.Gateway]())
	app.POST("/search", curd.ApiSearchWith[types.Gateway]([]*curd.Join{
		{"project", "project_id", "id", "name", "project"},
	}, "id", "name", "project_id", "disabled", "created"))
	app.GET("/list", curd.ApiList[types.Gateway]())
	app.POST("/create", curd.ApiCreateHook[types.Gateway](curd.GenerateID[types.Gateway](), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Gateway]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Gateway](nil, nil,
		"id", "name", "description", "project_id", "username", "password", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Gateway](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Gateway](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Gateway](false, nil, nil))
	app.GET("/export", export.ApiExport("gateway", "网关"))
	app.POST("/import", export.ApiImport("gateway"))
}
