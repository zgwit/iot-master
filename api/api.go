package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/config"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"net/http"
)

func catchError(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//runtime.Stack()
			//debug.Stack()
			switch err.(type) {
			case error:
				replyError(ctx, err.(error))
			case string:
				replyFail(ctx, err.(string))
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
		replyOk(ctx, &config.Config.Oem)
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
		replyOk(ctx, apps)
	})

	//修改配置
	app.GET("/config", loadConfig)
	app.POST("/config", saveConfig)

	//用户接口
	app.GET("/user/me", userMe)
	app.POST("/user/search", createCurdApiSearch[model.User]())
	app.GET("/user/list", createCurdApiList[model.User]())
	app.POST("/user/create", parseParamId, createCurdApiCreate[model.User](nil, nil))
	app.GET("/user/:id", parseParamId, createCurdApiGet[model.User]())
	app.POST("/user/:id", parseParamId, createCurdApiModify[model.User](nil, nil,
		"username", "name", "email", "disabled"))
	app.GET("/user/:id/delete", parseParamId, createCurdApiDelete[model.User](nil, nil))
	app.GET("/user/:id/password", parseParamId, userPassword)
	app.GET("/user/:id/enable", parseParamId, createCurdApiDisable[model.User](false, nil, nil))
	app.GET("/user/:id/disable", parseParamId, createCurdApiDisable[model.User](true, nil, nil))

	//产品接口
	app.POST("/product/search", createCurdApiSearch[model.Product]())
	app.GET("/product/list", createCurdApiList[model.Product]())
	app.POST("/product/create", createCurdApiCreate[model.Product](generateKey(8), nil))
	app.GET("/product/:id", parseParamStringId, createCurdApiGet[model.Product]())
	app.POST("/product/:id", parseParamStringId, createCurdApiModify[model.Product](nil, nil,
		"id", "name", "version", "desc", "properties", "functions", "events"))
	app.GET("/product/:id/delete", parseParamStringId, createCurdApiDelete[model.Product](nil, nil))

	//设备接口
	app.POST("/device/search", createCurdApiSearch[model.Device]())
	app.GET("/device/list", createCurdApiList[model.Device]())
	app.POST("/device/create", createCurdApiCreate[model.Device](generateKey(12), nil))
	app.GET("/device/:id", parseParamStringId, createCurdApiGet[model.Device]())
	app.POST("/device/:id", parseParamStringId, createCurdApiModify[model.Device](nil, nil,
		"id", "parent_id", "product_id", "type", "name", "desc", "username", "password", "disabled"))
	app.GET("/device/:id/delete", parseParamStringId, createCurdApiDelete[model.Device](nil, nil))

	app.GET("/device/:id/values", parseParamStringId, deviceValues)

	//报警日志
	app.POST("/alarm/search", createCurdApiSearch[model.Alarm]())
	app.GET("/alarm/list", createCurdApiList[model.Alarm]())
	app.GET("/alarm/:id", parseParamId, createCurdApiGet[model.Alarm]())
	app.GET("/alarm/:id/delete", parseParamId, createCurdApiDelete[model.Alarm](nil, nil))
	app.GET("/alarm/:id/read", parseParamId, alarmRead)

	//服务器接口
	app.POST("/server/search", createCurdApiSearch[model.Server]())
	app.GET("/server/list", createCurdApiList[model.Server]())
	app.POST("/server/create", createCurdApiCreate[model.Server](nil, nil))
	app.GET("/server/:id", parseParamId, createCurdApiGet[model.Server]())
	app.POST("/server/:id", parseParamId, createCurdApiModify[model.Server](nil, nil,
		"name", "type", "port", "desc", "disabled"))
	app.GET("/server/:id/delete", parseParamId, createCurdApiDelete[model.Server](nil, nil))

	//应用接口
	app.POST("/app/search", createCurdApiSearch[model.App]())
	app.GET("/app/list", createCurdApiList[model.App]())
	app.POST("/app/create", createCurdApiCreate[model.App](generateUUID, nil))
	app.GET("/app/:id", parseParamStringId, createCurdApiGet[model.App]())
	app.POST("/app/:id", parseParamStringId, createCurdApiModify[model.App](nil,
		nil, "id", "name", "type", "address", "desc", "disabled"))
	app.GET("/app/:id/delete", parseParamStringId, createCurdApiDelete[model.App](nil, nil))

	//插件接口
	app.POST("/plugin/search", createCurdApiSearch[model.Plugin]())
	app.GET("/plugin/list", createCurdApiList[model.Plugin]())
	app.POST("/plugin/create", createCurdApiCreate[model.Plugin](generateUUID, nil))
	app.GET("/plugin/:id", parseParamStringId, createCurdApiGet[model.Plugin]())
	app.POST("/plugin/:id", parseParamStringId, createCurdApiModify[model.Plugin](nil, nil,
		"id", "name", "version", "command", "dependencies"))
	app.GET("/plugin/:id/delete", parseParamStringId, createCurdApiDelete[model.Plugin](nil, nil))

	//系统接口
	app.GET("/system/cpu-info", cpuInfo)
	app.GET("/system/cpu", cpuStats)
	app.GET("/system/memory", memStats)
	app.GET("/system/disk", diskStats)
	app.GET("/system/machine", machineInfo)

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	app.Use(func(ctx *gin.Context) {
		replyFail(ctx, "Not found")
		ctx.Abort()
	})
}

func replyList(ctx *gin.Context, data interface{}, total int64) {
	ctx.JSON(http.StatusOK, gin.H{"data": data, "total": total})
}

func replyOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func replyFail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{"error": err})
}

func replyError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func nop(ctx *gin.Context) {
	ctx.String(http.StatusForbidden, "Unsupported")
}
