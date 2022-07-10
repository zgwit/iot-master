package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"iot-master/config"
	"iot-master/model"
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
	//检查Token
	token, has := ctx.GetQuery("token")
	if has {
		user, ok := tokens.Load(token)
		if ok {
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
			ctx.Abort()
		}
		return
	}

	//检查Session
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
		ctx.Abort()
	}
}

func RegisterRoutes(app *gin.RouterGroup) {
	//错误恢复，并返回至前端
	app.Use(catchError)

	app.GET("/info", info)

	app.POST("/login", login)
	app.POST("/auth", auth)

	//安装的接口
	if !config.Existing() {
		ins := app.Group("/install", func(ctx *gin.Context) {
			//仅限未安装的情况下调用
			if config.Existing() {
				replyFail(ctx, "已经安装过了")
				return
			}
			ctx.Next()
		})
		ins.POST("/base", installBase)
		ins.POST("/database", installDatabase)
		ins.POST("/history", installHistory)
		ins.GET("/system", installSystem)
	}

	//检查 session，必须登录
	app.Use(mustLogin)

	app.GET("/logout", logout)
	app.POST("/password", password)

	//修改配置
	app.GET("/config", loadConfig)
	app.POST("/config", saveConfig)

	//app.GET("/license", licenseDetail)
	//app.POST("/license", licenseUpdate)

	//用户接口
	app.GET("/user/me", userMe)
	app.POST("/user/list", createCurdApiList[model.User]())
	app.POST("/user/create", parseParamId, createCurdApiCreate[model.User](nil, nil))
	app.GET("/user/:id", parseParamId, createCurdApiGet[model.User]())
	app.POST("/user/:id", parseParamId, createCurdApiModify[model.User](nil, nil, "username", "nickname", "disabled"))
	app.GET("/user/:id/delete", parseParamId, createCurdApiDelete[model.User](nil, nil))
	app.GET("/user/:id/password", parseParamId, userPassword)
	app.GET("/user/:id/enable", parseParamId, createCurdApiDisable[model.User](false, nil, nil))
	app.GET("/user/:id/disable", parseParamId, createCurdApiDisable[model.User](true, nil, nil))

	//项目接口
	app.POST("/project/list", projectList)
	app.POST("/project/create", createCurdApiCreate[model.Project](nil, afterProjectCreate))
	app.GET("/project/:id", parseParamId, projectDetail)
	app.POST("/project/:id", parseParamId, createCurdApiModify[model.Project](nil, afterProjectUpdate,
		"name", "devices", "template_id", "hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"))
	app.GET("/project/:id/delete", parseParamId, createCurdApiDelete[model.Project](nil, afterProjectDelete))

	app.GET("/project/:id/start", parseParamId, projectStart)
	app.GET("/project/:id/stop", parseParamId, projectStop)
	app.GET("/project/:id/enable", parseParamId, createCurdApiDisable[model.Project](false, nil, afterProjectEnable))
	app.GET("/project/:id/disable", parseParamId, createCurdApiDisable[model.Project](true, nil, afterProjectDisable))
	app.GET("/project/:id/context", parseParamId, projectContext)
	app.POST("/project/:id/context", parseParamId, projectContextUpdate)
	app.GET("/project/:id/watch", parseParamId, projectWatch)

	//模板接口
	app.POST("/template/list", createCurdApiList[model.Template]())
	app.POST("/template/create", createCurdApiCreate[model.Template](generateUUID, nil))
	app.GET("/template/:id", parseParamStringId, createCurdApiGet[model.Template]())
	app.POST("/template/:id", parseParamStringId, createCurdApiModify[model.Template](nil, nil,
		"name", "info", "products", "hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"))
	app.GET("/template/:id/delete", parseParamStringId, createCurdApiDelete[model.Template](nil, nil))

	//设备接口
	app.POST("/device/list", deviceList)
	app.POST("/device/create", createCurdApiCreate[model.Device](nil, afterDeviceCreate))
	app.GET("/device/:id", parseParamId, deviceDetail)
	app.POST("/device/:id", parseParamId, createCurdApiModify[model.Device](nil, afterDeviceUpdate,
		"name", "tunnel_id", "product_id", "station",
		"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
		"context", "disabled",
	))
	app.GET("/device/:id/delete", parseParamId, createCurdApiDelete[model.Device](nil, afterDeviceDelete))

	app.GET("/device/:id/start", parseParamId, deviceStart)
	app.GET("/device/:id/stop", parseParamId, deviceStop)
	app.GET("/device/:id/enable", parseParamId, createCurdApiDisable[model.Device](false, nil, afterDeviceEnable))
	app.GET("/device/:id/disable", parseParamId, createCurdApiDisable[model.Device](true, nil, afterDeviceDisable))
	app.GET("/device/:id/context", parseParamId, deviceContext)
	app.POST("/device/:id/context", parseParamId, deviceContextUpdate)
	app.GET("/device/:id/watch", parseParamId, deviceWatch)
	app.GET("/device/:id/refresh", parseParamId, deviceRefresh)
	app.GET("/device/:id/refresh/:name", parseParamId, deviceRefreshPoint)
	app.POST("/device/:id/execute", parseParamId, deviceExecute)
	app.GET("/device/:id/value/:name/history", parseParamId, deviceValueHistory)

	//元件接口
	app.POST("/product/list", createCurdApiList[model.Product]())
	app.POST("/product/create", createCurdApiCreate[model.Product](generateUUID, nil))
	app.GET("/product/:id", parseParamStringId, createCurdApiGet[model.Product]())
	app.POST("/product/:id", parseParamStringId, createCurdApiModify[model.Product](nil, nil,
		"name", "manufacturer", "info",
		"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
		"context", "disabled",
	))
	app.GET("/product/:id/delete", parseParamStringId, createCurdApiDelete[model.Product](nil, nil))

	//组态
	app.POST("/hmi/list", createCurdApiList[model.Hmi]())
	app.POST("/hmi/create", createCurdApiCreate[model.Hmi](generateUUID, nil))
	app.GET("/hmi/:id", parseParamStringId, createCurdApiGet[model.Hmi]())
	app.POST("/hmi/:id", parseParamStringId, createCurdApiModify[model.Hmi](nil, nil, "name", "version"))
	app.GET("/hmi/:id/delete", parseParamStringId, createCurdApiDelete[model.Hmi](nil, nil))
	app.GET("/hmi/:id/manifest", hmiLoad)
	app.POST("/hmi/:id/manifest", hmiSave)

	app.GET("/hmi/:id/export", hmiExport)
	app.POST("/hmi/import") //zip

	//组件
	app.POST("/component/list", createCurdApiList[model.Component]())
	app.POST("/component/create", createCurdApiCreate[model.Component](generateUUID, nil))
	app.GET("/component/:id", parseParamStringId, createCurdApiGet[model.Component]())
	app.POST("/component/:id", parseParamStringId, createCurdApiModify[model.Component](nil, nil, "name", "group", "version"))
	app.GET("/component/:id/manifest", componentLoad)
	app.POST("/component/:id/manifest", componentSave)
	app.GET("/component/:id/delete", parseParamStringId, createCurdApiDelete[model.Component](nil, nil))

	app.GET("/component/:id/export", componentExport)
	app.POST("/component/import") //zip

	//服务器接口
	app.POST("/server/list", serverList)
	app.POST("/server/create", createCurdApiCreate[model.Server](nil, afterServerCreate))
	app.GET("/server/:id", parseParamId, serverDetail)
	app.POST("/server/:id", parseParamId, createCurdApiModify[model.Server](nil, afterServerUpdate,
		"name", "type", "addr",
		"retry", "register", "heartbeat", "protocol", "devices", "disabled"))
	app.GET("/server/:id/delete", parseParamId, createCurdApiDelete[model.Server](nil, afterServerDelete))

	app.GET("/server/:id/start", parseParamId, serverStart)
	app.GET("/server/:id/stop", parseParamId, serverStop)
	app.GET("/server/:id/enable", parseParamId, createCurdApiDisable[model.Server](false, nil, afterServerEnable))
	app.GET("/server/:id/disable", parseParamId, createCurdApiDisable[model.Server](true, nil, afterServerDisable))
	app.GET("/server/:id/watch", parseParamId, serverWatch)

	//通道接口
	app.POST("/tunnel/list", tunnelList)
	app.POST("/tunnel/create", createCurdApiCreate[model.Tunnel](nil, afterTunnelCreate))
	app.GET("/tunnel/:id", parseParamId, tunnelDetail)
	app.POST("/tunnel/:id", parseParamId, createCurdApiModify[model.Tunnel](nil, nil,
		"name", "type", "addr", "retry", "heartbeat", "serial", "protocol", "disabled"))
	app.GET("/tunnel/:id/delete", parseParamId, createCurdApiDelete[model.Tunnel](nil, afterTunnelDelete))
	app.GET("/tunnel/:id/start", parseParamId, tunnelStart)
	app.GET("/tunnel/:id/stop", parseParamId, tunnelClose)
	app.GET("/tunnel/:id/enable", parseParamId, createCurdApiDisable[model.Tunnel](false, nil, afterTunnelEnable))
	app.GET("/tunnel/:id/disable", parseParamId, createCurdApiDisable[model.Tunnel](true, nil, afterTunnelDisable))
	app.GET("/tunnel/:id/watch", parseParamId, tunnelWatch)

	//事件接口
	app.POST("/event/list", createCurdApiList[model.Event]())
	app.POST("/event/clear", eventClear)
	app.GET("/event/:id/delete", parseParamId, createCurdApiDelete[model.Event](nil, nil))

	//透传接口
	app.POST("/transfer/list", transferList)
	app.POST("/transfer/create", createCurdApiCreate[model.Transfer](nil, afterTransferCreate))
	app.GET("/transfer/:id", parseParamId, transferDetail)
	app.POST("/transfer/:id", parseParamId, createCurdApiModify[model.Transfer](nil, afterPipeUpdate,
		"name", "port", "tunnel_id", "disabled"))
	app.GET("/transfer/:id/delete", parseParamId, createCurdApiDelete[model.Transfer](nil, afterTransferDelete))

	app.GET("/transfer/:id/enable", parseParamId, createCurdApiDisable[model.Transfer](false, nil, afterTransferEnable))
	app.GET("/transfer/:id/disable", parseParamId, createCurdApiDisable[model.Transfer](true, nil, afterTransferDisable))
	app.GET("/transfer/:id/start", parseParamId, transferStart)
	app.GET("/transfer/:id/stop", parseParamId, transferStop)

	//摄像头接口
	app.POST("/camera/list", cameraList)
	app.POST("/camera/create", createCurdApiCreate[model.Camera](nil, afterCameraCreate))
	app.GET("/camera/:id", parseParamId, cameraDetail)
	app.POST("/camera/:id", parseParamId, createCurdApiModify[model.Camera](nil, afterCameraUpdate, "name", "url", "h264", "disabled"))
	app.GET("/camera/:id/delete", parseParamId, createCurdApiDelete[model.Camera](nil, afterCameraDelete))

	app.GET("/camera/:id/start", parseParamId, cameraStart)
	app.GET("/camera/:id/stop", parseParamId, cameraStop)
	app.GET("/camera/:id/enable", parseParamId, createCurdApiDisable[model.Camera](false, nil, afterCameraEnable))
	app.GET("/camera/:id/disable", parseParamId, createCurdApiDisable[model.Camera](true, nil, afterCameraDisable))

	//系统接口
	app.GET("/system/cpu-info", cpuInfo)
	app.GET("/system/cpu", cpuStats)
	app.GET("/system/memory", memStats)
	app.GET("/system/disk", diskStats)
	app.GET("/system/protocols", protocolList)
	app.GET("/system/protocol/:name", protocolDetail)
	app.GET("/system/serials", serialPortList)
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
