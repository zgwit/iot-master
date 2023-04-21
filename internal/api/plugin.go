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

// @Summary 导出插件
// @Schemes
// @Description 导出插件
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /plugin/export [get]
func noopPluginExport() {}

// @Summary 导入插件
// @Schemes
// @Description 导入插件
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回插件数量
// @Router /plugin/import [post]
func noopPluginImport() {}

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

	app.POST("/search", curd.ApiSearchHook[model.Plugin](func(datum []model.Plugin) error {
		for i := 0; i < len(datum); i++ {
			datum[i].Running = true
		}
		return nil
	}))

	app.GET("/list", curd.ApiListHook[model.Plugin](func(datum []model.Plugin) error {
		for i := 0; i < len(datum); i++ {
			datum[i].Running = true
		}
		return nil
	}))
	app.POST("/create", curd.ApiCreateHook[model.Plugin](curd.GenerateRandomId[model.Plugin](12), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGetHook[model.Plugin](func(m *model.Plugin) error {
		m.Running = true
		return nil
	}))

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.Plugin](nil, nil,
		"id", "name", "version", "command", "external", "username", "password", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.Plugin](nil, nil))

	app.GET("/export", curd.ApiExport[model.Plugin]("plugin"))
	app.POST("/import", curd.ApiImport[model.Plugin]())

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[model.Plugin](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[model.Plugin](false, nil, nil))

}
