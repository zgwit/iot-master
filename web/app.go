package web

import (
	"context"
	"embed"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/web/api"
	"log"
	"net/http"
	"time"
)

//go:embed www
var wwwFiles embed.FS

var server *http.Server

func Serve(cfg *Options) {
	if cfg == nil {
		cfg = DefaultOptions()
	}

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	//GIN初始化
	app := gin.Default()

	//启用session
	app.Use(sessions.Sessions("iot-master", memstore.NewStore([]byte("iot-master"))))

	//注册前端接口
	api.RegisterRoutes(app.Group("/api"))

	//前端静态文件
	//app.StaticFS("/", http.FS(wwwFiles))
	wwwFS := http.FS(wwwFiles)
	app.GET("/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		f, err := wwwFS.Open("www" + filepath)
		if err != nil {
			//默认首页
			f, err = wwwFS.Open("www/index.html")
			if err != nil {
				c.Next()
				return
			}
		}
		http.ServeContent(c.Writer, c.Request, filepath, time.Now(), f)
		_ = f.Close()
	})

	//监听HTTP
	//if err := app.Run(cfg.Addr); err != nil {
	//	log.Fatal("HTTP 服务启动错误", err)
	//}

	server = &http.Server{
		Addr:    cfg.Addr,
		Handler: app,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Web服务启动错误", err)
	}
}

func Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
