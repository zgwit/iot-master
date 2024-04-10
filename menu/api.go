package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
	"slices"
)

func init() {
	api.Register("GET", "/menu/:domain", menuGet)
}

// @Summary 获取菜单
// @Schemes
// @Description 获取菜单
// @Tags plugin
// @Param domain path string true "模块"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Menu] 返回插件信息
// @Router /menu/{domain} [get]
func menuGet(ctx *gin.Context) {
	domain := ctx.Param("domain")
	//TODO 获取用户权限，过滤菜单
	var ms []*Menu
	menus.Range(func(name string, m *Menu) bool {
		if slices.Contains(m.Domain, domain) {
			ms = append(ms, m)
		}
		return true
	})
	curd.OK(ctx, ms)
}
