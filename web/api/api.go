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
	app.POST("/login", login)
	app.GET("/logout", logout)

	//检查 session，必须登录
	app.Use(mustLogin)

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
	app.GET("/project/:id/watch", parseParamId, projectWatch)

	//模板接口
	modelTemplate := reflect.TypeOf(model.Template{})
	app.POST("/template/list", curdApiList(modelTemplate))
	app.POST("/template/create", curdApiCreate(modelTemplate, generateUUID, nil))
	app.GET("/template/:id", parseParamStringId, curdApiGet(modelTemplate))
	app.POST("/template/:id", parseParamStringId, curdApiModify(modelTemplate,
		[]string{"name", "version", "elements",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, nil))
	app.GET("/template/:id/delete", parseParamStringId, curdApiDelete(modelTemplate, nil, nil))

	//设备接口
	modelDevice := reflect.TypeOf(model.Device{})
	app.POST("/device/list", deviceList)
	app.POST("/device/create", curdApiCreate(modelDevice, nil, afterDeviceCreate))
	app.GET("/device/:id", parseParamId, deviceDetail)
	app.POST("/device/:id", parseParamId, curdApiModify(modelDevice,
		[]string{"name", "link_id", "element_id", "station",
			"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
			"context", "disabled"},
		nil, afterDeviceUpdate))
	app.GET("/device/:id/delete", parseParamId, curdApiDelete(modelDevice, nil, afterDeviceDelete))

	app.GET("/device/:id/start", parseParamId, deviceStart)
	app.GET("/device/:id/stop", parseParamId, deviceStop)
	app.GET("/device/:id/enable", parseParamId, curdApiDisable(modelDevice, false, nil, afterDeviceEnable))
	app.GET("/device/:id/disable", parseParamId, curdApiDisable(modelDevice, true, nil, afterDeviceDisable))
	app.GET("/device/:id/context", parseParamId, deviceContext)
	app.GET("/device/:id/watch", parseParamId, deviceWatch)
	app.GET("/device/:id/refresh", parseParamId, deviceRefresh)
	app.GET("/device/:id/refresh/:name", parseParamId, deviceRefreshPoint)
	app.POST("/device/:id/execute", parseParamId, deviceExecute)
	app.GET("/device/:id/value/:name/history", parseParamId, deviceValueHistory)

	//元件接口
	modelElement := reflect.TypeOf(model.Element{})
	app.POST("/element/list", curdApiList(modelElement))
	app.POST("/element/create", curdApiCreate(modelElement, generateUUID, nil))
	app.GET("/element/:id", parseParamStringId, curdApiGet(modelElement))
	app.POST("/element/:id", parseParamStringId, curdApiModify(modelElement,
		[]string{"name", "manufacturer", "version",
			"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
			"context", "disabled"},
		nil, nil))
	app.GET("/element/:id/delete", parseParamStringId, curdApiDelete(modelElement, nil, nil))

	//组态
	modelHMI := reflect.TypeOf(model.HMI{})
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

	//通道接口
	modelTunnel := reflect.TypeOf(model.Tunnel{})
	app.POST("/tunnel/list", tunnelList)
	app.POST("/tunnel/create", curdApiCreate(modelTunnel, nil, afterTunnelCreate))
	app.GET("/tunnel/:id", parseParamId, tunnelDetail)
	app.POST("/tunnel/:id", parseParamId, curdApiModify(modelTunnel,
		[]string{"name", "type", "addr",
			"retry", "register", "heartbeat", "serial", "protocol", "devices", "disabled"},
		nil, afterTunnelUpdate))
	app.GET("/tunnel/:id/delete", parseParamId, curdApiDelete(modelTunnel, nil, afterTunnelDelete))

	app.GET("/tunnel/:id/start", parseParamId, tunnelStart)
	app.GET("/tunnel/:id/stop", parseParamId, tunnelStop)
	app.GET("/tunnel/:id/enable", parseParamId, curdApiDisable(modelTunnel, false, nil, afterTunnelEnable))
	app.GET("/tunnel/:id/disable", parseParamId, curdApiDisable(modelTunnel, true, nil, afterTunnelDisable))
	app.GET("/tunnel/:id/watch", parseParamId, tunnelWatch)

	//连接接口
	modelLink := reflect.TypeOf(model.Link{})
	app.POST("/link/list", linkList)
	app.GET("/link/:id", parseParamId, linkDetail)
	app.POST("/link/:id", parseParamId, curdApiModify(modelLink,
		[]string{"name", "sn", "tunnel_id", "disabled"},
		nil, nil))
	app.GET("/link/:id/delete", parseParamId, curdApiDelete(modelLink, nil, afterLinkDelete))
	app.GET("/link/:id/stop", parseParamId, linkClose)
	app.GET("/link/:id/enable", parseParamId, curdApiDisable(modelLink, false, nil, nil))
	app.GET("/link/:id/disable", parseParamId, curdApiDisable(modelLink, true, nil, afterLinkDisable))
	app.GET("/link/:id/watch", parseParamId, linkWatch)

	//连接接口
	modelEvent := reflect.TypeOf(model.Event{})
	app.POST("/event/list", curdApiList(modelEvent))
	app.POST("/event/clear", eventClear)
	app.GET("/event/:id/delete", parseParamId, curdApiDelete(modelEvent, nil, nil))

	//连接接口
	modelPipe := reflect.TypeOf(model.Pipe{})
	app.POST("/pipe/list", pipeList)
	app.POST("/pipe/create", curdApiCreate(modelPipe, nil, afterPipeCreate))
	app.GET("/pipe/:id", parseParamId, pipeDetail)
	app.POST("/pipe/:id", parseParamId, curdApiModify(modelPipe,
		[]string{"name", "devices", "template_id",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, afterPipeUpdate))
	app.GET("/pipe/:id/delete", parseParamId, curdApiDelete(modelPipe, nil, afterPipeDelete))

	app.GET("/pipe/:id/enable", parseParamId, curdApiDisable(modelPipe, false, nil, afterPipeEnable))
	app.GET("/pipe/:id/disable", parseParamId, curdApiDisable(modelPipe, true, nil, afterPipeDisable))
	app.GET("/pipe/:id/start", parseParamId, pipeStart)
	app.GET("/pipe/:id/stop", parseParamId, pipeStop)

	//系统接口
	app.GET("/system/version", version)
	app.GET("/system/cpu-info", cpuInfo)
	app.GET("/system/cpu", cpuStats)
	app.GET("/system/memory", memStats)
	app.GET("/system/disk", diskStats)
	app.GET("/system/cron")
	app.GET("/system/protocols", protocolList)

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
