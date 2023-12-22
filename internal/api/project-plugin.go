package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
	"github.com/zgwit/iot-master/v4/types"
)

// @Summary 查询项目用户数量
// @Schemes
// @Description 查询项目用户数量
// @Tags project-plugin
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回项目用户数量
// @Router /project/plugin/count [post]
func noopProjectPluginCount() {}

// @Summary 查询项目用户
// @Schemes
// @Description 这里写描述 get project-plugins
// @Tags project-plugin
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.ProjectPlugin] 返回项目用户信息
// @Router /project/plugin/search [post]
func noopProjectPluginSearch() {}

// @Summary 查询项目用户
// @Schemes
// @Description 查询项目用户
// @Tags project-plugin
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.ProjectPlugin] 返回项目用户信息
// @Router /project/plugin/list [get]
func noopProjectPluginList() {}

// @Summary 创建项目用户
// @Schemes
// @Description 创建项目用户
// @Tags project-plugin
// @Param search body types.ProjectPlugin true "项目用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectPlugin] 返回项目用户信息
// @Router /project/plugin/create [post]
func noopProjectPluginCreate() {}

// @Summary 修改项目用户
// @Schemes
// @Description 修改项目用户
// @Tags project-plugin
// @Param id path int true "项目用户ID"
// @Param project-plugin body types.ProjectPlugin true "项目用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectPlugin] 返回项目用户信息
// @Router /project/plugin/{id} [post]
func noopProjectPluginUpdate() {}

// @Summary 获取项目用户
// @Schemes
// @Description 获取项目用户
// @Tags project-plugin
// @Param id path int true "项目用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectPlugin] 返回项目用户信息
// @Router /project/plugin/{id} [get]
func noopProjectPluginGet() {}

// @Summary 删除项目用户
// @Schemes
// @Description 删除项目用户
// @Tags project-plugin
// @Param id path int true "项目用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectPlugin] 返回项目用户信息
// @Router /project/plugin/{id}/delete [get]
func noopProjectPluginDelete() {}

func projectPluginRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.ProjectPlugin]())

	app.POST("/search", curd.ApiSearch[types.ProjectPlugin]())

	app.GET("/list", curd.ApiList[types.ProjectPlugin]())

	app.POST("/create", curd.ApiCreate[types.ProjectPlugin]())

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.ProjectPlugin]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.ProjectPlugin](nil, nil,
		"id", "project_id", "plugin_id", "disabled"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.ProjectPlugin](nil, nil))

}
