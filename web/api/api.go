package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/model"
	"net/http"
	"reflect"
)

func mustLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		//TODO 检查OAuth2返回的code，进一步获取用户信息，放置到session中

		ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
		ctx.Abort()
	}
}

func RegisterRoutes(app *gin.RouterGroup) {
	//错误恢复，并返回至前端
	app.Use(func(ctx *gin.Context) {
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

	})

	app.POST("/login", login)

	//检查 session，必须登录
	app.Use(mustLogin)

	app.GET("/logout", logout)
	app.GET("/password", password)

	//用户接口
	modelUser := reflect.TypeOf(model.User{})
	app.GET("/user/me", userMe)
	app.POST("/user/list", curdApiList(modelUser))
	app.POST("/user/create", parseParamId, curdApiCreate(modelUser, nil, nil))
	app.GET("/user/:id", parseParamId, curdApiGet(modelUser))
	app.POST("/user/:id", parseParamId, curdApiModify(modelUser, []string{"username", "nickname", "disabled"}, nil, nil))
	app.GET("/user/:id/delete", parseParamId, curdApiDelete(modelUser, nil, nil))
	app.GET("/user/:id/password", parseParamId, userPassword)
	app.GET("/user/:id/enable", parseParamId, curdApiDisable(modelUser, false, nil, nil))
	app.GET("/user/:id/disable", parseParamId, curdApiDisable(modelUser, true, nil, nil))

	//项目接口
	modelProject := reflect.TypeOf(model.Project{})
	app.POST("/project/list", projectList)
	app.POST("/project/create", curdApiCreate(modelProject, nil, afterProjectCreate))
	app.GET("/project/:id", parseParamId, projectDetail)
	app.POST("/project/:id", parseParamId, curdApiModify(modelProject,
		[]string{"name", "devices", "template_id",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, afterProjectUpdate))
	app.GET("/project/:id/delete", parseParamId, curdApiDelete(modelProject, nil, afterProjectDelete))

	app.GET("/project/:id/start", parseParamId, projectStart)
	app.GET("/project/:id/stop", parseParamId, projectStop)
	app.GET("/project/:id/enable", parseParamId, curdApiDisable(modelProject, false, nil, afterProjectEnable))
	app.GET("/project/:id/disable", parseParamId, curdApiDisable(modelProject, true, nil, afterProjectDisable))
	app.GET("/project/:id/context", parseParamId, projectContext)
	app.POST("/project/:id/context", parseParamId, projectContextUpdate)
	app.GET("/project/:id/watch", parseParamId, projectWatch)

	//模板接口
	modelTemplate := reflect.TypeOf(model.Template{})
	app.POST("/template/list", curdApiList(modelTemplate))
	app.POST("/template/create", curdApiCreate(modelTemplate, generateUUID, nil))
	app.GET("/template/:id", parseParamStringId, curdApiGet(modelTemplate))
	app.POST("/template/:id", parseParamStringId, curdApiModify(modelTemplate,
		[]string{"name", "version", "products",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, nil))
	app.GET("/template/:id/delete", parseParamStringId, curdApiDelete(modelTemplate, nil, nil))

	//设备接口
	modelDevice := reflect.TypeOf(model.Device{})
	app.POST("/device/list", deviceList)
	app.POST("/device/create", curdApiCreate(modelDevice, nil, afterDeviceCreate))
	app.GET("/device/:id", parseParamId, deviceDetail)
	app.POST("/device/:id", parseParamId, curdApiModify(modelDevice,
		[]string{"name", "tunnel_id", "product_id", "station",
			"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
			"context", "disabled"},
		nil, afterDeviceUpdate))
	app.GET("/device/:id/delete", parseParamId, curdApiDelete(modelDevice, nil, afterDeviceDelete))

	app.GET("/device/:id/start", parseParamId, deviceStart)
	app.GET("/device/:id/stop", parseParamId, deviceStop)
	app.GET("/device/:id/enable", parseParamId, curdApiDisable(modelDevice, false, nil, afterDeviceEnable))
	app.GET("/device/:id/disable", parseParamId, curdApiDisable(modelDevice, true, nil, afterDeviceDisable))
	app.GET("/device/:id/context", parseParamId, deviceContext)
	app.POST("/device/:id/context", parseParamId, deviceContextUpdate)
	app.GET("/device/:id/watch", parseParamId, deviceWatch)
	app.GET("/device/:id/refresh", parseParamId, deviceRefresh)
	app.GET("/device/:id/refresh/:name", parseParamId, deviceRefreshPoint)
	app.POST("/device/:id/execute", parseParamId, deviceExecute)
	app.GET("/device/:id/value/:name/history", parseParamId, deviceValueHistory)

	//元件接口
	modelProduct := reflect.TypeOf(model.Product{})
	app.POST("/product/list", curdApiList(modelProduct))
	app.POST("/product/create", curdApiCreate(modelProduct, generateUUID, nil))
	app.GET("/product/:id", parseParamStringId, curdApiGet(modelProduct))
	app.POST("/product/:id", parseParamStringId, curdApiModify(modelProduct,
		[]string{"name", "manufacturer", "version",
			"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
			"context", "disabled"},
		nil, nil))
	app.GET("/product/:id/delete", parseParamStringId, curdApiDelete(modelProduct, nil, nil))

	//组态
	modelHMI := reflect.TypeOf(model.Hmi{})
	app.POST("/hmi/list", curdApiList(modelHMI))
	app.POST("/hmi/create", curdApiCreate(modelHMI, generateUUID, nil))
	app.GET("/hmi/:id", parseParamStringId, curdApiGet(modelHMI))
	app.POST("/hmi/:id", parseParamStringId, curdApiModify(modelHMI,
		[]string{"name", "width", "height", "snap", "entities"},
		nil, nil))
	app.GET("/hmi/:id/delete", parseParamStringId, curdApiDelete(modelHMI, nil, nil))
	app.GET("/hmi/:id/export")

	//组态的附件
	app.GET("/hmi/:id/attachment/*name", hmiAttachmentRead)
	app.POST("/hmi/:id/attachment/*name", hmiAttachmentUpload)
	app.PATCH("/hmi/:id/attachment/*name", hmiAttachmentRename)
	app.DELETE("/hmi/:id/attachment/*name", hmiAttachmentDelete)

	//服务器接口
	modelServer := reflect.TypeOf(model.Server{})
	app.POST("/server/list", serverList)
	app.POST("/server/create", curdApiCreate(modelServer, nil, afterServerCreate))
	app.GET("/server/:id", parseParamId, serverDetail)
	app.POST("/server/:id", parseParamId, curdApiModify(modelServer,
		[]string{"name", "type", "addr",
			"retry", "register", "heartbeat", "protocol", "devices", "disabled"},
		nil, afterServerUpdate))
	app.GET("/server/:id/delete", parseParamId, curdApiDelete(modelServer, nil, afterServerDelete))

	app.GET("/server/:id/start", parseParamId, serverStart)
	app.GET("/server/:id/stop", parseParamId, serverStop)
	app.GET("/server/:id/enable", parseParamId, curdApiDisable(modelServer, false, nil, afterServerEnable))
	app.GET("/server/:id/disable", parseParamId, curdApiDisable(modelServer, true, nil, afterServerDisable))
	app.GET("/server/:id/watch", parseParamId, serverWatch)

	//通道接口
	modelTunnel := reflect.TypeOf(model.Tunnel{})
	app.POST("/tunnel/list", tunnelList)
	app.POST("/tunnel/create", curdApiCreate(modelTunnel, nil, afterTunnelCreate))
	app.GET("/tunnel/:id", parseParamId, tunnelDetail)
	app.POST("/tunnel/:id", parseParamId, curdApiModify(modelTunnel,
		[]string{"name", "type", "addr", "retry", "heartbeat", "serial", "protocol", "disabled"},
		nil, nil))
	app.GET("/tunnel/:id/delete", parseParamId, curdApiDelete(modelTunnel, nil, afterTunnelDelete))
	app.GET("/tunnel/:id/start", parseParamId, tunnelStart)
	app.GET("/tunnel/:id/stop", parseParamId, tunnelClose)
	app.GET("/tunnel/:id/enable", parseParamId, curdApiDisable(modelTunnel, false, nil, nil))
	app.GET("/tunnel/:id/disable", parseParamId, curdApiDisable(modelTunnel, true, nil, afterTunnelDisable))
	app.GET("/tunnel/:id/watch", parseParamId, tunnelWatch)

	//事件接口
	modelEvent := reflect.TypeOf(model.Event{})
	app.POST("/event/list", curdApiList(modelEvent))
	app.POST("/event/clear", eventClear)
	app.GET("/event/:id/delete", parseParamId, curdApiDelete(modelEvent, nil, nil))

	//透传接口
	modelTransfer := reflect.TypeOf(model.Transfer{})
	app.POST("/transfer/list", transferList)
	app.POST("/transfer/create", curdApiCreate(modelTransfer, nil, afterTransferCreate))
	app.GET("/transfer/:id", parseParamId, transferDetail)
	app.POST("/transfer/:id", parseParamId, curdApiModify(modelTransfer,
		[]string{"name", "port", "tunnel_id", "disabled"},
		nil, afterPipeUpdate))
	app.GET("/transfer/:id/delete", parseParamId, curdApiDelete(modelTransfer, nil, afterTransferDelete))

	app.GET("/transfer/:id/enable", parseParamId, curdApiDisable(modelTransfer, false, nil, afterTransferEnable))
	app.GET("/transfer/:id/disable", parseParamId, curdApiDisable(modelTransfer, true, nil, afterTransferDisable))
	app.GET("/transfer/:id/start", parseParamId, transferStart)
	app.GET("/transfer/:id/stop", parseParamId, transferStop)

	//摄像头接口
	modelCamera := reflect.TypeOf(model.Camera{})
	app.POST("/camera/list", cameraList)
	app.POST("/camera/create", curdApiCreate(modelCamera, nil, afterCameraCreate))
	app.GET("/camera/:id", parseParamId, cameraDetail)
	app.POST("/camera/:id", parseParamId, curdApiModify(modelCamera, []string{"name", "url", "disabled"}, nil, afterCameraUpdate))
	app.GET("/camera/:id/delete", parseParamId, curdApiDelete(modelCamera, nil, afterCameraDelete))

	app.GET("/camera/:id/start", parseParamId, cameraStart)
	app.GET("/camera/:id/stop", parseParamId, cameraStop)
	app.GET("/camera/:id/enable", parseParamId, curdApiDisable(modelCamera, false, nil, afterCameraEnable))
	app.GET("/camera/:id/disable", parseParamId, curdApiDisable(modelCamera, true, nil, afterCameraDisable))

	//系统接口
	app.GET("/system/version", version)
	app.GET("/system/cpu-info", cpuInfo)
	app.GET("/system/cpu", cpuStats)
	app.GET("/system/memory", memStats)
	app.GET("/system/disk", diskStats)
	app.GET("/system/cron")
	app.GET("/system/protocols", protocolList)
	app.GET("/system/protocol/:name", protocolDetail)

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
