package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/web"
	"github.com/zgwit/iot-master/v4/web/attach"
	"github.com/zgwit/iot-master/v4/web/curd"
	"net/http"
)

type API struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
}

var apis []*API

func Register(method, path string, handlers ...gin.HandlerFunc) {
	apis = append(apis, &API{method, path, handlers})
}

func catchError(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//runtime.Stack()
			//debug.Stack()
			switch err.(type) {
			case error:
				curd.Error(ctx, err.(error))
			case string:
				curd.Fail(ctx, err.(string))
			default:
				ctx.JSON(http.StatusOK, gin.H{"error": err})
			}
			//TODO 这里好像又继续了
		}
	}()
	ctx.Next()

	//TODO 内容如果为空，返回404

}

func mustLogin(ctx *gin.Context) {
	token := ctx.Request.URL.Query().Get("token")
	if token == "" {
		token = ctx.Request.Header.Get("Authorization")
		if token != "" {
			//此处需要去掉 Bearer
			token = token[7:]
		}
	}

	if token != "" {
		claims, err := web.JwtVerify(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		} else {
			ctx.Set("user", claims.Id) //与session统一
			ctx.Next()
		}
		return
	}

	//检查Session
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
	}
}

func RegisterRoutes(router *gin.RouterGroup) {
	//错误恢复，并返回至前端
	router.Use(catchError)

	router.GET("/auth", auth)
	router.POST("/login", login)

	router.GET("/oem", oem)
	router.GET("/info", info)

	//检查 session，必须登录
	router.Use(mustLogin)

	router.GET("/logout", logout)
	router.POST("/password", password)

	//注册接口
	for _, a := range apis {
		router.Handle(a.Method, a.Path, a.Handlers...)
	}

	//OEM
	oemRouter(router.Group("/oem"))

	//配置文件
	settingRouter(router.Group("/setting"))

	//注册子接口
	userRouter(router.Group("/user"))

	pluginRouter(router.Group("/plugin"))

	productRouter(router.Group("/product"))
	attach.Routers(router.Group("/product/:id/attach"), "product", "id")

	productVersionRouter(router.Group("/product/:id/version"))
	attach.Routers(router.Group("/product/:id/version/:version/attach"), "product", "id", "version")

	projectRouter(router.Group("/project"))
	projectUserRouter(router.Group("/project/:id/user"))
	attach.Routers(router.Group("/project/:id/attach"), "project", "id")

	spaceRouter(router.Group("/space"))
	spaceDeviceRouter(router.Group("/space/:id/device"))
	attach.Routers(router.Group("/space/:id/attach"), "space", "id")

	gatewayRouter(router.Group("/gateway"))

	deviceRouter(router.Group("/device"))

	alarmRouter(router.Group("/alarm"))

	backupRouter(router.Group("/backup"))

	//附件管理
	attach.Routers(router.Group("/attach"), "attach")

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	router.Use(func(ctx *gin.Context) {
		curd.Fail(ctx, "Not found")
		ctx.Abort()
	})
}
