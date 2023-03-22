package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"mime"
	"net/http"
	"path"
	"time"
)

func init() {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		log.Error(err)
	}
}

func CreateEngine(cfg Options) *gin.Engine {
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

	//app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}

func RegisterFS(app *gin.Engine, fs http.FileSystem) {
	tm := time.Now()
	app.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			//支持前端框架的无“#”路由
			fn := path.Join("www", c.Request.URL.Path) //删除查询参数
			f, err := fs.Open(fn)
			if err == nil {
				defer f.Close()
				stat, err := f.Stat()
				if err != nil {
					c.Next() //500错误
					return
				}
				if !stat.IsDir() {
					http.ServeContent(c.Writer, c.Request, fn, tm, f)
					return
				}
			}

			//默认首页
			f, err = fs.Open("www/index.html")
			if err != nil {
				c.Next()
				return
			}
			defer f.Close()

			fn += ".html" //避免DetectContentType
			http.ServeContent(c.Writer, c.Request, fn, tm, f)
		}
	})

}
