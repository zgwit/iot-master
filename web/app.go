package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/zgwit/iot-master/v3/internal"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var proxies internal.Map[httputil.ReverseProxy]

func appProxy(ctx *gin.Context) {
	p := proxies.Load(ctx.Param("app"))

	if p == nil {
		app := internal.Applications.Load(ctx.Param("app"))
		if app == nil {
			_ = ctx.Error(errors.New("应用未注册"))
			return
		}
		u, err := url.Parse(app.Address)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		//创建反向代理服务
		p = &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = u.Scheme
				req.URL.Host = u.Host
				//req.URL.Path = u.Path
				//设置User-Agent(
				if _, ok := req.Header["User-Agent"]; !ok {
					// explicitly disable User-Agent so it's not set to default value
					req.Header.Set("User-Agent", "")
				}
			},
		}
		//支持unixsock加速
		if app.Type == "unix" {
			p.Transport = &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", app.Address)
			}}
		}
		proxies.Store(app.Id, p)
	}

	p.ServeHTTP(ctx.Writer, ctx.Request)

	ctx.Abort()
}
