package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/model"
)

// @Summary 查询角色数量
// @Schemes
// @Description 查询角色数量
// @Tags role
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回角色数量
// @Router /role/count [post]
func noopRoleCount() {}

// @Summary 查询角色
// @Schemes
// @Description 这里写描述 get roles
// @Tags role
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Role] 返回角色信息
// @Router /role/search [post]
func noopRoleSearch() {}

// @Summary 查询角色
// @Schemes
// @Description 查询角色
// @Tags role
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Role] 返回角色信息
// @Router /role/list [get]
func noopRoleList() {}

// @Summary 创建角色
// @Schemes
// @Description 创建角色
// @Tags role
// @Param search body model.Role true "角色信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Role] 返回角色信息
// @Router /role/create [post]
func noopRoleCreate() {}

// @Summary 修改角色
// @Schemes
// @Description 修改角色
// @Tags role
// @Param id path string true "角色ID"
// @Param role body model.Role true "角色信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Role] 返回角色信息
// @Router /role/{id} [post]
func noopRoleUpdate() {}

// @Summary 获取角色
// @Schemes
// @Description 获取角色
// @Tags role
// @Param id path string true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Role] 返回角色信息
// @Router /role/{id} [get]
func noopRoleGet() {}

// @Summary 删除角色
// @Schemes
// @Description 删除角色
// @Tags role
// @Param id path string true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Role] 返回角色信息
// @Router /role/{id}/delete [get]
func noopRoleDelete() {}

// @Summary 导出角色
// @Schemes
// @Description 导出角色
// @Tags role
// @Accept json
// @Produce octet-stream
// @Success 200 {object} curd.ReplyList[model.Role] 返回压缩包
// @Router /role/export [get]
func noopRoleExport() {}

// @Summary 导入角色
// @Schemes
// @Description 导入角色
// @Tags role
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回角色数量
// @Router /role/import [post]
func noopRoleImport() {}

func roleRouter(app *gin.RouterGroup) {

	app.POST("/count", curd2.ApiCount[model.Role]())

	app.POST("/search", curd2.ApiSearch[model.Role]())

	app.GET("/list", curd2.ApiList[model.Role]())

	app.POST("/create", curd2.ParseParamStringId, curd2.ApiCreateHook[model.Role](
		curd2.GenerateRandomId[model.Role](9), nil))

	app.GET("/:id", curd2.ParseParamStringId, curd2.ApiGet[model.Role]())

	app.POST("/:id", curd2.ParseParamStringId, curd2.ApiUpdateHook[model.Role](nil, nil,
		"id", "name", "privileges"))

	app.GET("/:id/delete", curd2.ParseParamStringId, curd2.ApiDeleteHook[model.Role](nil, nil))

	app.GET("/export", curd2.ApiExport("role", "角色"))

	app.POST("/import", curd2.ApiImport("role"))

}
