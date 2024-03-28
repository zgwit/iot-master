package plugin

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {

	api.Register("GET", "/plugin/list", func(ctx *gin.Context) {
		curd.OK(ctx, GetPlugins())
	})

	api.Register("GET", "/plugin/:id/manifest", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := Get(ctx.GetString("id"))
		if p == nil {
			curd.Fail(ctx, "插件未加载")
			return
		}
		curd.OK(ctx, p.Manifest)
	})

	api.Register("GET", "/plugin/menus/:entry", menus)
}

// @Summary 查询插件
// @Schemes
// @Description 查询插件
// @Tags plugin
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Manifest] 返回插件信息
// @Router /plugin/list [get]
func noopPluginList() {}

// @Summary 获取插件详情
// @Schemes
// @Description 获取插件详情
// @Tags plugin
// @Param id path int true "插件ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Manifest] 返回插件信息
// @Router /plugin/{id}/manifest [get]
func noopPluginManifestGet() {}

// @Summary 获取插件菜单
// @Schemes
// @Description 获取插件菜单
// @Tags plugin
// @Param entry path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Menu] 返回插件信息
// @Router /plugin/menus/{entry} [get]
func menus(ctx *gin.Context) {
	entry := ctx.Param("entry")

	var menus []*Menu
	for _, p := range GetPlugins() {
		if p.Menus != nil {
			if en, ok := p.Menus[entry]; ok {
				menus = append(menus, en)
			}
		}
	}
	curd.OK(ctx, menus)
}
