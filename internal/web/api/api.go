package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v2/model"
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

	//项目接口
	app.POST("/project/search", createCurdApiSearch[model.Project]())
	app.GET("/project/list", createCurdApiList[model.Project]())
	app.POST("/project/create", createCurdApiCreate[model.Project](generateUUID, afterProjectCreate))
	app.GET("/project/:id", parseParamStringId, createCurdApiGet[model.Project]())
	app.POST("/project/:id", parseParamStringId, createCurdApiModify[model.Project](nil, afterProjectUpdate,
		"name", "devices", "interface_id", "disabled"))
	app.GET("/project/:id/delete", parseParamStringId, createCurdApiDelete[model.Project](nil, afterProjectDelete))

	app.GET("/project/:id/values", parseParamStringId, projectValues)
	app.POST("/project/:id/assign", parseParamStringId, projectAssign)
	//app.GET("/project/:id/targets", parseParamStringId, projectTargets)

	//设备接口
	app.POST("/device/search", createCurdApiSearch[model.Device]())
	app.GET("/device/list", createCurdApiList[model.Device]())
	app.POST("/device/create", createCurdApiCreate[model.Device](generateUUID, afterDeviceCreate))
	app.GET("/device/:id", parseParamStringId, createCurdApiGet[model.Device]())
	app.POST("/device/:id", parseParamStringId, createCurdApiModify[model.Device](nil, afterDeviceUpdate,
		"name", "tunnel_id", "product_id", "interface_id", "station", "disabled"))
	app.GET("/device/:id/delete", parseParamStringId, createCurdApiDelete[model.Device](afterDeviceDelete, nil))

	app.GET("/device/:id/values", parseParamStringId, deviceValues)
	app.POST("/device/:id/assign", parseParamStringId, deviceAssign)
	app.GET("/device/:id/refresh", parseParamStringId, deviceRefresh)

	//元件接口
	app.POST("/product/search", createCurdApiSearch[model.Product]())
	app.GET("/product/list", createCurdApiList[model.Product]())
	app.POST("/product/create", createCurdApiCreate[model.Product](generateUUID, nil))
	app.GET("/product/:id", parseParamStringId, createCurdApiGet[model.Product]())
	app.POST("/product/:id", parseParamStringId, createCurdApiModify[model.Product](nil, nil,
		"name", "manufacturer", "version", "protocol", "points", "pollers", "disabled"))
	app.GET("/product/:id/delete", parseParamStringId, createCurdApiDelete[model.Product](nil, nil))

	//网关接口
	app.POST("/gateway/search", createCurdApiSearch[model.Gateway]())
	app.GET("/gateway/list", createCurdApiList[model.Gateway]())
	app.POST("/gateway/create", createCurdApiCreate[model.Gateway](generateUUID, nil))
	app.GET("/gateway/:id", parseParamStringId, createCurdApiGet[model.Gateway]())
	app.POST("/gateway/:id", parseParamStringId, createCurdApiModify[model.Gateway](nil, nil,
		"name", "type", "version"))
	app.GET("/gateway/:id/delete", parseParamStringId, createCurdApiDelete[model.Gateway](nil, nil))

	//网关的通道和服务
	app.GET("/gateway/:id/list/tunnel", parseParamStringId, createCurdApiListWithId[model.Tunnel]("gateway_id"))
	app.GET("/gateway/:id/list/server", parseParamStringId, createCurdApiListWithId[model.Server]("gateway_id"))

	//服务器接口
	app.POST("/server/search", createCurdApiSearch[model.Server]())
	app.GET("/server/list", createCurdApiList[model.Server]())
	app.POST("/server/create", createCurdApiCreate[model.Server](generateUUID, afterServerCreate))
	app.GET("/server/:id", parseParamStringId, createCurdApiGet[model.Server]())
	app.POST("/server/:id", parseParamStringId, createCurdApiModify[model.Server](nil, afterServerUpdate,
		"name", "type", "addr",
		"retry", "register", "heartbeat", "protocol", "devices", "disabled"))
	app.GET("/server/:id/delete", parseParamStringId, createCurdApiDelete[model.Server](nil, afterServerDelete))

	//服务的通道
	app.GET("/server/:id/list/tunnel", parseParamStringId, createCurdApiListWithId[model.Tunnel]("server_id"))

	//通道接口
	app.POST("/tunnel/search", createCurdApiSearch[model.Tunnel]())
	app.GET("/tunnel/list", createCurdApiList[model.Tunnel]())
	app.POST("/tunnel/create", createCurdApiCreate[model.Tunnel](generateUUID, afterTunnelCreate))
	app.GET("/tunnel/:id", parseParamStringId, createCurdApiGet[model.Tunnel]())
	app.POST("/tunnel/:id", parseParamStringId, createCurdApiModify[model.Tunnel](nil, afterTunnelUpdate,
		"name", "type", "addr", "retry", "serial", "protocol", "disabled"))
	app.GET("/tunnel/:id/delete", parseParamStringId, createCurdApiDelete[model.Tunnel](nil, afterTunnelDelete))

	//通道的设备
	app.GET("/tunnel/:id/list/device", parseParamStringId, createCurdApiListWithId[model.Device]("tunnel_id"))

	//界面接口
	app.POST("/interface/search", createCurdApiSearch[model.Interface]())
	app.GET("/interface/list", createCurdApiList[model.Interface]())
	app.POST("/interface/create", createCurdApiCreate[model.Interface](generateUUID, nil))
	app.GET("/interface/:id", parseParamStringId, createCurdApiGet[model.Interface]())
	app.POST("/interface/:id", parseParamStringId, createCurdApiModify[model.Interface](nil, nil,
		"name"))
	app.GET("/interface/:id/delete", parseParamStringId, createCurdApiDelete[model.Interface](nil, nil))

	//插件接口
	app.POST("/plugin/search", createCurdApiSearch[model.Plugin]())
	app.GET("/plugin/list", createCurdApiList[model.Plugin]())
	app.POST("/plugin/create", createCurdApiCreate[model.Plugin](generateUUID, nil))
	app.GET("/plugin/:id", parseParamStringId, createCurdApiGet[model.Plugin]())
	app.POST("/plugin/:id", parseParamStringId, createCurdApiModify[model.Plugin](nil, nil,
		"name"))
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
