package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/config"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"net/http"
)

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
		claims, err := jwtVerify(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		}
		ctx.Set("user", claims.Id) //与session统一
		ctx.Next()
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

func RegisterRoutes(app *gin.RouterGroup) {
	//错误恢复，并返回至前端
	app.Use(catchError)

	app.GET("/oem", func(ctx *gin.Context) {
		curd.OK(ctx, &config.Config.Oem)
	})

	app.GET("/info", info)

	app.GET("/auth", auth)
	app.POST("/login", login)

	//检查 session，必须登录
	app.Use(mustLogin)

	app.GET("/logout", logout)
	app.POST("/password", password)

	app.GET("/apps", func(ctx *gin.Context) {
		apps := make([]*model.App, 0)
		internal.Applications.Range(func(name string, app *model.App) bool {
			if !app.Hidden {
				apps = append(apps, app)
			}
			return true
		})
		curd.OK(ctx, apps)
	})

	//修改配置
	app.GET("/config", loadConfig)
	app.POST("/config", saveConfig)

	//注册子接口
	userRouter(app.Group("/user"))
	productRouter(app.Group("/product"))
	deviceRouter(app.Group("/device"))
	groupRouter(app.Group("/group"))
	alarmRouter(app.Group("/alarm"))
	serverRouter(app.Group("/server"))
	appRouter(app.Group("/app"))
	pluginRouter(app.Group("/plugin"))
	systemRouter(app.Group("/system"))

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	app.Use(func(ctx *gin.Context) {
		curd.Fail(ctx, "Not found")
		ctx.Abort()
	})
}
