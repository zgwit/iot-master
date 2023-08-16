package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/plugin"
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
// @Tags plugin
// @Accept json
// @Produce octet-stream
// @Router /plugin/export [get]
func noopPluginExport() {}

// @Summary 导入插件
// @Schemes
// @Description 导入插件
// @Tags plugin
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

// @Summary 启动插件
// @Schemes
// @Description 启动插件
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/{id}/start [get]
func noopPluginStart() {}

// @Summary 停止插件
// @Schemes
// @Description 停止插件
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Plugin] 返回插件信息
// @Router /plugin/{id}/stop [get]
func noopPluginStop() {}

func pluginRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearchHook[model.Plugin](func(datum []*model.Plugin) error {
		for _, v := range datum {
			p := plugin.Get(v.Id)
			if p != nil {
				v.Running = p.Running
			}
		}
		return nil
	}))

	app.GET("/list", curd.ApiListHook[model.Plugin](func(datum []*model.Plugin) error {
		for _, v := range datum {
			p := plugin.Get(v.Id)
			if p != nil {
				v.Running = p.Running
			}
		}
		return nil
	}))
	app.POST("/create", curd.ApiCreateHook[model.Plugin](curd.GenerateRandomId[model.Plugin](12), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGetHook[model.Plugin](func(m *model.Plugin) error {
		p := plugin.Get(m.Id)
		if p != nil {
			m.Running = p.Running
		}
		return nil
	}))

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.Plugin](nil, nil,
		"id", "name", "version", "command", "external", "username", "password", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.Plugin](nil, nil))

	app.GET("/export", curd.ApiExport("plugin", "插件"))

	app.POST("/import", curd.ApiImport("plugin"))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[model.Plugin](true, nil, func(id any) error {
		p := plugin.Get(id.(string))
		if p == nil {
			return errors.New("插件未加载")
		}
		err := p.Close()
		if err != nil {
			return err
		}
		return nil
	}))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[model.Plugin](false, nil, func(id any) error {
		return plugin.Load(id.(string))
	}))

	app.GET(":id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		err := plugin.Load(ctx.GetString("id"))
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	app.GET(":id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := plugin.Get(ctx.GetString("id"))
		if p == nil {
			curd.Fail(ctx, "插件未加载")
			return
		}
		err := p.Close()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

}
