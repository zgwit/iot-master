package user

import (
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

// @Summary 查询角色
// @Schemes
// @Description 查询角色
// @Tags role
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Role] 返回角色信息
// @Router /role/search [post]
func noopRoleSearch() {}

// @Summary 查询角色
// @Schemes
// @Description 查询角色
// @Tags role
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Role] 返回角色信息
// @Router /role/list [get]
func noopRoleList() {}

// @Summary 创建角色
// @Schemes
// @Description 创建角色
// @Tags role
// @Param search body Role true "角色信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Role] 返回角色信息
// @Router /role/create [post]
func noopRoleCreate() {}

// @Summary 修改角色
// @Schemes
// @Description 修改角色
// @Tags role
// @Param id path int true "角色ID"
// @Param role body Role true "角色信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Role] 返回角色信息
// @Router /role/{id} [post]
func noopRoleUpdate() {}

// @Summary 删除角色
// @Schemes
// @Description 删除角色
// @Tags role
// @Param id path int true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Role] 返回角色信息
// @Router /role/{id}/delete [get]
func noopRoleDelete() {}

// @Summary 启用角色
// @Schemes
// @Description 启用角色
// @Tags role
// @Param id path int true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Role] 返回角色信息
// @Router /role/{id}/enable [get]
func noopRoleEnable() {}

// @Summary 禁用角色
// @Schemes
// @Description 禁用角色
// @Tags role
// @Param id path int true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Role] 返回角色信息
// @Router /role/{id}/disable [get]
func noopRoleDisable() {}

func init() {

	api.Register("POST", "/count", curd.ApiCount[Role]())

	api.Register("POST", "/search", curd.ApiSearch[Role]())

	api.Register("GET", "/list", curd.ApiList[Role]())

	api.Register("POST", "/create", curd.ApiCreateHook[Role](curd.GenerateID[Role](), nil))

	api.Register("GET", "/:id", curd.ParseParamStringId, curd.ApiGet[Role]())

	api.Register("POST", "/:id", curd.ParseParamStringId, curd.ApiUpdate[Role]("id", "name", "privileges", "description", "disabled"))

	api.Register("GET", "/:id/delete", curd.ParseParamStringId, curd.ApiDelete[Role]())

	api.Register("GET", ":id/disable", curd.ParseParamStringId, curd.ApiDisable[Role](true))

	api.Register("GET", "/:id/enable", curd.ParseParamStringId, curd.ApiDisable[Role](false))

}
