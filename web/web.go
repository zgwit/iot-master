package web

import (
	"context"
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/zgwit/iot-master/v3/api"
	"github.com/zgwit/iot-master/v3/args"
	"github.com/zgwit/iot-master/v3/config"
	"github.com/zgwit/iot-master/v3/docs"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/swagger-files"
	"mime"
	"net/http"
	"path"
	"strings"
	"time"
)

func init() {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		log.Error(err)
	}
}

//go:embed all:www
var wwwFiles embed.FS

var server *http.Server

func Serve(cfg config.Web) {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	//GIN初始化
	//app := gin.Default()
	app := gin.New()
	app.Use(gin.Recovery())

	if cfg.Debug {
		app.Use(gin.Logger())
	}

	//跨域问题
	if cfg.Cors {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = true
		c.AllowCredentials = true
		app.Use(cors.New(c))
	}

	//启用session
	app.Use(sessions.Sessions("iot-master", cookie.NewStore([]byte("iot-master"))))

	//开启压缩
	if cfg.Gzip {
		app.Use(gzip.Gzip(gzip.DefaultCompression)) //gzip.WithExcludedPathsRegexs([]string{".*"})
	}

	//注册前端接口
	api.RegisterRoutes(app.Group("/api"))

	docs.SwaggerInfo.BasePath = "/api"
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//app.Any("/app/:id/*path", appProxy)

	//使用$前缀区分插件
	app.Any("/$:app/*path", appProxy)

	//前端静态文件
	//app.StaticFS("/www", http.FS(wwwFiles))

	wwwFS := http.FS(wwwFiles)
	app.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			//支持前端框架的无“#”路由
			fn := path.Join("www", c.Request.URL.Path) //删除查询参数
			f, err := wwwFS.Open(fn)
			if err == nil {
				defer f.Close()
				stat, err := f.Stat()
				if err != nil {
					c.Next() //500错误
					return
				}
				if !stat.IsDir() {
					http.ServeContent(c.Writer, c.Request, fn, args.Time, f)
					return
				}
			}

			//默认首页
			f, err = wwwFS.Open("www/index.html")
			if err != nil {
				c.Next()
				return
			}
			defer f.Close()

			fn += ".html" //避免DetectContentType
			http.ServeContent(c.Writer, c.Request, fn, args.Time, f)
		}
	})

	//监听HTTP
	//if err := app.Run(cfg.Addr); err != nil {
	//	log.Fatal("HTTP 服务启动错误", err)
	//}

	log.Println("Web服务启动", cfg.Addr)
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

func resolvePort(addr string) string {
	if strings.IndexByte(addr, ':') == -1 {
		return ":" + addr
	}
	return addr
}
