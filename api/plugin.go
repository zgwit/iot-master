package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/plugin"
	"github.com/zgwit/iot-master/v4/web/curd"
)

// @Summary 查询插件
// @Schemes
// @Description 查询插件
// @Tags plugin
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[plugin.Manifest] 返回插件信息
// @Router /plugin/list [get]
func noopPluginList() {}

// @Summary 获取插件详情
// @Schemes
// @Description 获取插件详情
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[plugin.Manifest] 返回插件信息
// @Router /plugin/{id}/manifest [get]
func noopPluginManifestGet() {}

// @Summary 获取插件菜单
// @Schemes
// @Description 获取插件菜单
// @Tags plugin
// @Param entry path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]plugin.Menu] 返回插件信息
// @Router /plugin/menus/{entry} [get]
func menus(ctx *gin.Context) {
	entry := ctx.Param("entry")

	var menus []*plugin.Menu
	for _, p := range plugin.GetPlugins() {
		if p.Menus != nil {
			if en, ok := p.Menus[entry]; ok {
				menus = append(menus, en)
			}
		}
	}
	curd.OK(ctx, menus)
}

// @Summary 获取插件入口
// @Schemes
// @Description 获取插件入口
// @Tags plugin
// @Param entry path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]plugin.Page] 返回插件信息
// @Router /plugin/pages/{entry} [get]
func pages(ctx *gin.Context) {
	entry := ctx.Param("entry")

	var entries []*plugin.Page
	for _, p := range plugin.GetPlugins() {
		if p.Pages != nil {
			if en, ok := p.Pages[entry]; ok {
				entries = append(entries, en)
			}
		}
	}
	curd.OK(ctx, entries)
}

func pluginRouter(app *gin.RouterGroup) {

	app.GET("/list", func(ctx *gin.Context) {
		curd.OK(ctx, plugin.GetPlugins())
	})

	app.GET("/:id/manifest", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := plugin.Get(ctx.GetString("id"))
		if p == nil {
			curd.Fail(ctx, "插件未加载")
			return
		}
		curd.OK(ctx, p.Manifest)
	})

	app.GET("/menus/:entry", menus)

	app.GET("/pages/:entry", pages)
}
