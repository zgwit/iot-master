package broker

import (
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {
	api.Register("POST", "/gateway/count", curd.ApiCount[Gateway]())
	api.Register("POST", "/gateway/search", curd.ApiSearch[Gateway]("id", "name", "disabled", "created"))
	api.Register("GET", "/gateway/list", curd.ApiList[Gateway]())
	api.Register("POST", "/gateway/create", curd.ApiCreateHook[Gateway](curd.GenerateID[Gateway](), nil))
	api.Register("GET", "/gateway/:id", curd.ParseParamStringId, curd.ApiGet[Gateway]())
	api.Register("POST", "/gateway/:id", curd.ParseParamStringId, curd.ApiUpdateHook[Gateway](nil, nil,
		"id", "name", "description", "username", "password", "disabled"))
	api.Register("GET", "/gateway/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Gateway](nil, nil))
	api.Register("GET", "/gateway/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[Gateway](true, nil, nil))
	api.Register("GET", "/gateway/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[Gateway](false, nil, nil))
}

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
// @Success 200 {object} curd.ReplyList[Gateway] 返回网关信息
// @Router /gateway/search [post]
func noopGatewaySearch() {}

// @Summary 查询网关
// @Schemes
// @Description 查询网关
// @Tags gateway
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Gateway] 返回网关信息
// @Router /gateway/list [get]
func noopGatewayList() {}

// @Summary 创建网关
// @Schemes
// @Description 创建网关
// @Tags gateway
// @Param search body Gateway true "网关信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Gateway] 返回网关信息
// @Router /gateway/create [post]
func noopGatewayCreate() {}

// @Summary 修改网关
// @Schemes
// @Description 修改网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Param gateway body Gateway true "网关信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Gateway] 返回网关信息
// @Router /gateway/{id} [post]
func noopGatewayUpdate() {}

// @Summary 获取网关
// @Schemes
// @Description 获取网关
// @Tags gateway
// @Param id path string true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Gateway] 返回网关信息
// @Router /gateway/{id} [get]
func noopGatewayGet() {}

// @Summary 删除网关
// @Schemes
// @Description 删除网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Gateway] 返回网关信息
// @Router /gateway/{id}/delete [get]
func noopGatewayDelete() {}

// @Summary 启用网关
// @Schemes
// @Description 启用网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Gateway] 返回网关信息
// @Router /gateway/{id}/enable [get]
func noopGatewayEnable() {}

// @Summary 禁用网关
// @Schemes
// @Description 禁用网关
// @Tags gateway
// @Param id path int true "网关ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Gateway] 返回网关信息
// @Router /gateway/{id}/disable [get]
func noopGatewayDisable() {}
