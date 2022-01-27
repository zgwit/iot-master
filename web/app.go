package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"iot-master/conf"
	"iot-master/web/api"
	wwwFiles "iot-master/web/www"
	"log"
	"net/http"
)

func Serve() {
	if !conf.Config.Web.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	//GIN初始化
	app := gin.Default()

	//启用session
	app.Use(sessions.Sessions("iot-master", memstore.NewStore([]byte("iot-master"))))

	//注册前端接口
	api.RegisterRoutes(app.Group("/api"))

	//前端静态文件
	//app.GET("/*any", func(c *gin.Context) {
	app.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			//支持前端框架的无“#”路由
			if c.Request.RequestURI == "/" {
				c.Request.URL.Path = "index.html"
			} else if _, err := wwwFiles.FS.Stat(wwwFiles.CTX, c.Request.RequestURI); err != nil {
				c.Request.URL.Path = "index.html"
			}
			//TODO 如果未登录，则跳转SysAdmin OAuth2自动授权页面

			//文件失效期已经在Handler中处理
			wwwFiles.Handler.ServeHTTP(c.Writer, c.Request)
		}
	})

	//监听HTTP
	if err := app.Run(conf.Config.Web.Addr); err != nil {
		log.Fatal("HTTP 服务启动错误", err)
	}
}