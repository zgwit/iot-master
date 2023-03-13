package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询插件
// @Schemes
// @Description 查询插件
// @Tags plugin
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Plugin] 返回插件信息
// @Router /plugin/search [post]
func noopPluginSearch() {}

// @Summary 查询插件
// @Schemes
// @Description 查询插件
// @Tags plugin
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Plugin] 返回插件信息
// @Router /plugin/list [get]
func noopPluginList() {}

// @Summary 创建插件
// @Schemes
// @Description 创建插件
// @Tags plugin
// @Param search body model.Plugin true "插件信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/create [post]
func noopPluginCreate() {}

// @Summary 修改插件
// @Schemes
// @Description 修改插件
// @Tags plugin
// @Param id path int true "插件ID"
// @Param plugin body model.Plugin true "插件信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/{id} [post]
func noopPluginUpdate() {}

// @Summary 删除插件
// @Schemes
// @Description 删除插件
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/{id}/delete [get]
func noopPluginDelete() {}

// @Summary 启用插件
// @Schemes
// @Description 启用插件
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/{id}/enable [get]
func noopPluginEnable() {}

// @Summary 禁用插件
// @Schemes
// @Description 禁用插件
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/{id}/disable [get]
func noopPluginDisable() {}

func pluginRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Plugin]())
	app.GET("/list", curd.ApiList[model.Plugin]())
	app.POST("/create", curd.ApiCreate[model.Plugin](curd.GenerateUuidKey, nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Plugin]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Plugin](nil, nil,
		"id", "name", "version", "command", "dependencies"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Plugin](nil, nil))
}
