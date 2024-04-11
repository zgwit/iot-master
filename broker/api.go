package broker

import (
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {
	api.Register("POST", "/broker/count", curd.ApiCount[Broker]())
	api.Register("POST", "/broker/search", curd.ApiSearch[Broker]("id", "name", "port", "disabled", "created"))
	api.Register("GET", "/broker/list", curd.ApiList[Broker]())
	api.Register("POST", "/broker/create", curd.ApiCreateHook[Broker](curd.GenerateID[Broker](), nil))
	api.Register("GET", "/broker/:id", curd.ParseParamStringId, curd.ApiGet[Broker]())
	api.Register("POST", "/broker/:id", curd.ParseParamStringId, curd.ApiUpdate[Broker]())
	api.Register("GET", "/broker/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Broker](nil, nil))
	api.Register("GET", "/broker/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[Broker](true, nil, nil))
	api.Register("GET", "/broker/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[Broker](false, nil, nil))
}

// @Summary 查询MQTT服务器数量
// @Schemes
// @Description 查询MQTT服务器数量
// @Tags broker
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回MQTT服务器数量
// @Router /broker/count [post]
func noopBrokerCount() {}

// @Summary 查询MQTT服务器
// @Schemes
// @Description 查询MQTT服务器
// @Tags broker
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Broker] 返回MQTT服务器信息
// @Router /broker/search [post]
func noopBrokerSearch() {}

// @Summary 查询MQTT服务器
// @Schemes
// @Description 查询MQTT服务器
// @Tags broker
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Broker] 返回MQTT服务器信息
// @Router /broker/list [get]
func noopBrokerList() {}

// @Summary 创建MQTT服务器
// @Schemes
// @Description 创建MQTT服务器
// @Tags broker
// @Param search body Broker true "MQTT服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Broker] 返回MQTT服务器信息
// @Router /broker/create [post]
func noopBrokerCreate() {}

// @Summary 修改MQTT服务器
// @Schemes
// @Description 修改MQTT服务器
// @Tags broker
// @Param id path int true "MQTT服务器ID"
// @Param broker body Broker true "MQTT服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Broker] 返回MQTT服务器信息
// @Router /broker/{id} [post]
func noopBrokerUpdate() {}

// @Summary 获取MQTT服务器
// @Schemes
// @Description 获取MQTT服务器
// @Tags broker
// @Param id path string true "MQTT服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Broker] 返回MQTT服务器信息
// @Router /broker/{id} [get]
func noopBrokerGet() {}

// @Summary 删除MQTT服务器
// @Schemes
// @Description 删除MQTT服务器
// @Tags broker
// @Param id path int true "MQTT服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Broker] 返回MQTT服务器信息
// @Router /broker/{id}/delete [get]
func noopBrokerDelete() {}

// @Summary 启用MQTT服务器
// @Schemes
// @Description 启用MQTT服务器
// @Tags broker
// @Param id path int true "MQTT服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Broker] 返回MQTT服务器信息
// @Router /broker/{id}/enable [get]
func noopBrokerEnable() {}

// @Summary 禁用MQTT服务器
// @Schemes
// @Description 禁用MQTT服务器
// @Tags broker
// @Param id path int true "MQTT服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Broker] 返回MQTT服务器信息
// @Router /broker/{id}/disable [get]
func noopBrokerDisable() {}
