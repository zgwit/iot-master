package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"net/http"
	"path"
	"time"
)

type Engine struct {
	*gin.Engine
}

func CreateEngine() *Engine {
	if !options.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	//GIN初始化
	//app := gin.Default()
	app := gin.New()
	app.Use(gin.Recovery())

	if options.Debug {
		app.Use(gin.Logger())
	}

	//跨域问题
	if options.Cors {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = true
		c.AllowCredentials = true
		app.Use(cors.New(c))
	}

	//启用session
	app.Use(sessions.Sessions("iot-master", cookie.NewStore([]byte("iot-master"))))

	//开启压缩
	if options.Gzip {
		app.Use(gzip.Gzip(gzip.DefaultCompression)) //gzip.WithExcludedPathsRegexs([]string{".*"})
	}

	//app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Engine{app}
}

func (app *Engine) FileSystem() *FileSystem {
	var fs FileSystem
	tm := time.Now()
	app.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			f, err := fs.Open(c.Request.URL.Path)
			if err == nil {
				defer f.Close()
				stat, err := f.Stat()
				if err != nil {
					c.Next() //500错误
					return
				}
				if !stat.IsDir() {
					fn := c.Request.URL.Path + ".html" //避免DetectContentType
					http.ServeContent(c.Writer, c.Request, fn, tm, f)
					return
				}
			}
		}
	})
	return &fs
}

func (app *Engine) RegisterFS(fs http.FileSystem, prefix, index string) {
	tm := time.Now()
	app.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			//支持前端框架的无“#”路由
			fn := path.Join(prefix, c.Request.URL.Path) //删除查询参数
			//fn := c.Request.URL.Path
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
			fn = path.Join(prefix, index) //删除查询参数
			f, err = fs.Open(fn)
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

func (app *Engine) Serve() {
	log.Info("Web服务启动 ", options.Addr)
	err := app.Run(options.Addr)
	if err != nil {
		log.Fatal("HTTP 服务启动错误", err)
	}
}
