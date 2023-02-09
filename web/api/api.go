package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal/config"
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

	//修改配置
	app.GET("/config", loadConfig)
	app.POST("/config", saveConfig)

	//用户接口
	app.GET("/user/me", userMe)
	app.POST("/user/search", createCurdApiSearch[model.User]())
	app.GET("/user/list", createCurdApiList[model.User]())
	app.POST("/user/create", parseParamId, createCurdApiCreate[model.User](nil, nil))
	app.GET("/user/:id", parseParamId, createCurdApiGet[model.User]())
	app.POST("/user/:id", parseParamId, createCurdApiModify[model.User](nil, nil, "username", "nickname", "disabled"))
	app.GET("/user/:id/delete", parseParamId, createCurdApiDelete[model.User](nil, nil))
	app.GET("/user/:id/password", parseParamId, userPassword)
	app.GET("/user/:id/enable", parseParamId, createCurdApiDisable[model.User](false, nil, nil))
	app.GET("/user/:id/disable", parseParamId, createCurdApiDisable[model.User](true, nil, nil))

	//网关接口
	app.POST("/gateway/search", createCurdApiSearch[model.Gateway]())
	app.GET("/gateway/list", createCurdApiList[model.Gateway]())
	app.POST("/gateway/create", createCurdApiCreate[model.Gateway](generateUUID, nil))
	app.GET("/gateway/:id", parseParamStringId, createCurdApiGet[model.Gateway]())
	app.POST("/gateway/:id", parseParamStringId, createCurdApiModify[model.Gateway](nil, nil,
		"name", "desc", "username", "password", "client_id", "disabled"))
	app.GET("/gateway/:id/delete", parseParamStringId, createCurdApiDelete[model.Gateway](afterDeviceDelete, nil))

	app.GET("/gateway/:id/properties", parseParamStringId, gatewayProperties)

	//设备接口
	app.POST("/device/search", createCurdApiSearch[model.Device]())
	app.GET("/device/list", createCurdApiList[model.Device]())
	app.POST("/device/create", createCurdApiCreate[model.Device](generateUUID, afterDeviceCreate))
	app.GET("/device/:id", parseParamStringId, createCurdApiGet[model.Device]())
	app.POST("/device/:id", parseParamStringId, createCurdApiModify[model.Device](nil, afterDeviceUpdate, "name", "tunnel_id", "model_id", "interface_id", "station", "disabled"))
	app.GET("/device/:id/delete", parseParamStringId, createCurdApiDelete[model.Device](afterDeviceDelete, nil))

	app.GET("/device/:id/properties", parseParamStringId, deviceProperties)
	//app.POST("/device/:id/assign", parseParamStringId, deviceAssign)
	//app.GET("/device/:id/refresh", parseParamStringId, deviceRefresh)

	//元件接口
	app.POST("/model/search", createCurdApiSearch[model.Model]())
	app.GET("/model/list", createCurdApiList[model.Model]())
	app.POST("/model/create", createCurdApiCreate[model.Model](generateUUID, nil))
	app.GET("/model/:id", parseParamStringId, createCurdApiGet[model.Model]())
	app.POST("/model/:id", parseParamStringId, createCurdApiModify[model.Model](nil, nil,
		"name", "manufacturer", "version", "protocol", "points", "pollers", "disabled"))
	app.GET("/model/:id/delete", parseParamStringId, createCurdApiDelete[model.Model](nil, nil))

	//服务器接口
	app.POST("/server/search", createCurdApiSearch[model.Server]())
	app.GET("/server/list", createCurdApiList[model.Server]())
	app.POST("/server/create", createCurdApiCreate[model.Server](nil, nil))
	app.GET("/server/:id", parseParamId, createCurdApiGet[model.Server]())
	app.POST("/server/:id", parseParamId, createCurdApiModify[model.Server](nil, nil,
		"name", "type", "address", "desc", "disabled"))
	app.GET("/server/:id/delete", parseParamId, createCurdApiDelete[model.Server](nil, nil))

	//服务器接口
	app.POST("/service/search", createCurdApiSearch[model.App]())
	app.GET("/service/list", createCurdApiList[model.App]())
	app.POST("/service/create", createCurdApiCreate[model.App](generateUUID, nil))
	app.GET("/service/:id", parseParamStringId, createCurdApiGet[model.App]())
	app.POST("/service/:id", parseParamStringId, createCurdApiModify[model.App](nil,
		nil, "name", "type", "address", "desc", "disabled"))
	app.GET("/service/:id/delete", parseParamStringId, createCurdApiDelete[model.App](nil, nil))

	//插件接口
	app.POST("/plugin/search", createCurdApiSearch[model.Plugin]())
	app.GET("/plugin/list", createCurdApiList[model.Plugin]())
	app.POST("/plugin/create", createCurdApiCreate[model.Plugin](generateUUID, nil))
	app.GET("/plugin/:id", parseParamStringId, createCurdApiGet[model.Plugin]())
	app.POST("/plugin/:id", parseParamStringId, createCurdApiModify[model.Plugin](nil, nil, "name"))
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
