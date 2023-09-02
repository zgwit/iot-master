package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/model"
	"github.com/zgwit/iot-master/v4/web"
	"net/http/httputil"
)

var proxies lib.Map[httputil.ReverseProxy]

func ProxyApp(ctx *gin.Context) {
	p := proxies.Load(ctx.Param("app"))

	if p == nil {
		app := Applications.Load(ctx.Param("app"))
		if app == nil {
			_ = ctx.Error(errors.New("应用未注册"))
			return
		}

		rp, err := web.CreateReverseProxy(app.Type, app.Address)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		p = rp
		proxies.Store(app.Id, rp)
	}

	p.ServeHTTP(ctx.Writer, ctx.Request)

	ctx.Abort()
}

var Applications lib.Map[model.App]

func Register(app *model.App) {
	Applications.Store(app.Id, app)
}
