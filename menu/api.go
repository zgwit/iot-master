package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"slices"
)

func init() {
	api.Register("GET", "/menu/:domain", menuGet)
}

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
	api.OK(ctx, ms)
}
