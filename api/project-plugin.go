package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/curd"
	"xorm.io/xorm/schemas"
)

// @Summary 项目插件列表
// @Schemes
// @Description 项目插件列表
// @Tags project-plugin
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.ProjectPlugin] 返回项目插件信息
// @Router /project/{id}/plugin [get]
func projectPluginList(ctx *gin.Context) {
	var pds []types.ProjectPlugin
	err := db.Engine.Where("project_id=?", ctx.Param("id")).Find(&pds)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, pds)
}

// @Summary 绑定项目插件
// @Schemes
// @Description 绑定项目插件
// @Tags project-plugin
// @Param id path int true "项目ID"
// @Param plugin path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/plugin/{plugin}/bind [get]
func projectPluginBind(ctx *gin.Context) {
	pd := types.ProjectPlugin{
		ProjectId: ctx.Param("id"),
		PluginId:  ctx.Param("plugin"),
	}
	_, err := db.Engine.InsertOne(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 删除项目插件
// @Schemes
// @Description 删除项目插件
// @Tags project-plugin
// @Param id path int true "项目ID"
// @Param plugin path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/plugin/{plugin}/unbind [get]
func projectPluginUnbind(ctx *gin.Context) {
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("plugin")}).Delete(new(types.ProjectPlugin))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 禁用项目插件
// @Schemes
// @Description 禁用项目插件
// @Tags project-plugin
// @Param id path int true "项目ID"
// @Param plugin path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/plugin/{plugin}/disable [get]
func projectPluginDisable(ctx *gin.Context) {
	pd := types.ProjectPlugin{Disabled: true}
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("plugin")}).Cols("disabled").Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 启用项目插件
// @Schemes
// @Description 启用项目插件
// @Tags project-plugin
// @Param id path int true "项目ID"
// @Param plugin path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/plugin/{plugin}/enable [get]
func projectPluginEnable(ctx *gin.Context) {
	pd := types.ProjectPlugin{Disabled: false}
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("plugin")}).Cols("disabled").Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 修改项目插件
// @Schemes
// @Description 修改项目插件
// @Tags project-plugin
// @Param id path int true "项目ID"
// @Param plugin path int true "插件ID"
// @Param project-plugin body types.ProjectPlugin true "项目插件信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/plugin/{plugin} [post]
func projectPluginUpdate(ctx *gin.Context) {
	var pd types.ProjectPlugin
	err := ctx.ShouldBindJSON(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	_, err = db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("plugin")}).
		Cols("plugin_id", "disabled").
		Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func projectPluginRouter(app *gin.RouterGroup) {
	app.GET("", projectPluginList)
	app.GET("/:plugin/bind", projectPluginBind)
	app.GET("/:plugin/unbind", projectPluginUnbind)
	app.GET("/:plugin/disable", projectPluginDisable)
	app.GET("/:plugin/enable", projectPluginEnable)
	app.POST("/:plugin", projectPluginUpdate)
}
