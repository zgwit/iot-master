package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/model"
	"net/http"
	"reflect"
	"xorm.io/xorm"
)

type paramSearchEx struct {
	Skip     int                    `form:"skip" json:"skip"`
	Limit    int                    `form:"limit" json:"limit"`
	Sort     map[string]int         `form:"sort" json:"sort"`
	Filters  map[string]interface{} `form:"filter" json:"filter"`
	Keywords map[string]string      `form:"keyword" json:"keyword"`
}

func (body *paramSearchEx) toQuery() *xorm.Session {
	if body.Limit < 1 {
		body.Limit = 20
	}
	op := db.Engine.Limit(body.Limit, body.Skip)

	for k, v := range body.Filters {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			ll := len(v.([]interface{}))
			if ll > 0 {
				if ll == 1 {
					op.And(k+"=?", v.([]interface{})[0])
				} else {
					op.In(k, v)
				}
			}
		} else {
			if v != nil {
				op.And(k+"=?", v)
			}
		}
	}

	for k, v := range body.Keywords {
		if v != "" {
			op.And(k+" like", "%"+v+"%")
		}
	}

	if len(body.Sort) > 0 {
		for k, v := range body.Sort {
			if v > 0 {
				op.Asc(k)
			} else {
				op.Desc(k)
			}
		}
	} else {
		op.Desc("id")
	}

	return op
}

type paramId struct {
	Id int64 `uri:"id"`
}
type paramStringId struct {
	Id string `uri:"id"`
}

type WatchMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

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

func parseParamId(ctx *gin.Context) {
	var pid paramId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
}

func parseParamStringId(ctx *gin.Context) {
	var pid paramStringId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
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
	app.POST("/user/create", curdApiCreate(modelUser, nil, nil))
	app.GET("/user/:id", curdApiGet(modelUser))
	app.POST("/user/:id", curdApiModify(modelUser, []string{"username", "nickname", "disabled"}, nil, nil))
	app.GET("/user/:id/delete", curdApiDelete(modelUser, nil, nil))
	app.GET("/user/:id/password", parseParamId, userPassword)
	app.GET("/user/:id/enable", curdApiDisable(modelUser, false, nil, nil))
	app.GET("/user/:id/disable", curdApiDisable(modelUser, true, nil, nil))

	//项目接口
	modelProject := reflect.TypeOf(model.Project{})
	app.POST("/project/list", curdApiList(modelProject))
	app.POST("/project/create", curdApiCreate(modelProject, nil, afterProjectCreate))
	app.GET("/project/:id", curdApiGet(modelProject))
	app.POST("/project/:id", curdApiModify(modelProject,
		[]string{"name", "devices", "template_id",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, afterProjectUpdate))
	app.GET("/project/:id/delete", curdApiDelete(modelProject, nil, afterProjectDelete))

	app.GET("/project/:id/start", parseParamId, projectStart)
	app.GET("/project/:id/stop", parseParamId, projectStop)
	app.GET("/project/:id/enable", curdApiDisable(modelProject, false, nil, afterProjectEnable))
	app.GET("/project/:id/disable", curdApiDisable(modelProject, true, nil, afterProjectDisable))
	app.GET("/project/:id/context", parseParamId, projectContext)
	app.GET("/project/:id/watch", parseParamId, projectWatch)

	//模板接口
	modelTemplate := reflect.TypeOf(model.Template{})
	app.POST("/template/list", curdApiList(modelTemplate))
	app.POST("/template/create", curdApiCreate(modelTemplate, nil, nil))
	app.GET("/template/:id", curdApiGet(modelTemplate))
	app.POST("/template/:id", curdApiModify(modelTemplate,
		[]string{"name", "version", "elements",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, nil))
	app.GET("/template/:id/delete", curdApiDelete(modelTemplate, nil, nil))

	//设备接口
	modelDevice := reflect.TypeOf(model.Device{})
	app.POST("/device/list", deviceList)
	app.POST("/device/create", curdApiCreate(modelDevice, nil, afterDeviceCreate))
	app.GET("/device/:id", curdApiGet(modelDevice))
	app.POST("/device/:id", curdApiModify(modelDevice,
		[]string{"name", "link_id", "element_id", "station",
			"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
			"context", "disabled"},
		nil, afterDeviceUpdate))
	app.GET("/device/:id/delete", curdApiDelete(modelDevice, nil, afterDeviceDelete))

	app.GET("/device/:id/start", parseParamId, deviceStart)
	app.GET("/device/:id/stop", parseParamId, deviceStop)
	app.GET("/device/:id/enable", curdApiDisable(modelDevice, false, nil, afterDeviceEnable))
	app.GET("/device/:id/disable", curdApiDisable(modelDevice, true, nil, afterDeviceDisable))
	app.GET("/device/:id/context", parseParamId, deviceContext)
	app.GET("/device/:id/watch", parseParamId, deviceWatch)
	app.GET("/device/:id/refresh", parseParamId, deviceRefresh)
	app.GET("/device/:id/refresh/:name", parseParamId, deviceRefreshPoint)
	app.POST("/device/:id/execute", parseParamId, deviceExecute)
	app.GET("/device/:id/value/:name/history", parseParamId, deviceValueHistory)

	//元件接口
	modelElement := reflect.TypeOf(model.Element{})
	app.POST("/element/list", curdApiList(modelElement))
	app.POST("/element/create", curdApiCreate(modelElement, nil, nil))
	app.GET("/element/:id", curdApiGet(modelElement))
	app.POST("/element/:id", curdApiModify(modelElement,
		[]string{"name", "manufacturer", "version",
			"hmi", "tags", "points", "pollers", "calculators", "alarms", "commands",
			"context", "disabled"},
		nil, nil))
	app.GET("/element/:id/delete", curdApiDelete(modelElement, nil, nil))

	//组态
	modelHMI := reflect.TypeOf(model.HMI{})
	app.POST("/hmi/list", curdApiList(modelHMI))
	app.POST("/hmi/create", curdApiCreate(modelHMI, nil, nil))
	app.GET("/hmi/:id", curdApiGet(modelHMI))
	app.POST("/hmi/:id", curdApiModify(modelHMI,
		[]string{"name", "width", "height", "snap", "entities"},
		nil, nil))
	app.GET("/hmi/:id/delete", curdApiDelete(modelHMI, nil, nil))
	app.GET("/hmi/:id/export")

	//组态的附件
	app.GET("/hmi/:id/attachment/*name", parseParamId, hmiAttachmentRead)
	app.POST("/hmi/:id/attachment/*name", parseParamId, hmiAttachmentUpload)
	app.PATCH("/hmi/:id/attachment/*name", parseParamId, hmiAttachmentRename)
	app.DELETE("/hmi/:id/attachment/*name", parseParamId, hmiAttachmentDelete)

	//通道接口
	modelTunnel := reflect.TypeOf(model.Tunnel{})
	app.POST("/tunnel/list", curdApiList(modelTunnel))
	app.POST("/tunnel/create", curdApiCreate(modelTunnel, nil, afterTunnelCreate))
	app.GET("/tunnel/:id", tunnelDetail)
	app.POST("/tunnel/:id", curdApiModify(modelTunnel,
		[]string{"name", "devices", "template_id",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, afterTunnelUpdate))
	app.GET("/tunnel/:id/delete", curdApiDelete(modelTunnel, nil, afterTunnelDelete))

	app.GET("/tunnel/:id/start", parseParamId, tunnelStart)
	app.GET("/tunnel/:id/stop", parseParamId, tunnelStop)
	app.GET("/tunnel/:id/enable", curdApiDisable(modelTunnel, false, nil, afterTunnelEnable))
	app.GET("/tunnel/:id/disable", curdApiDisable(modelTunnel, true, nil, afterTunnelDisable))
	app.GET("/tunnel/:id/watch", parseParamId, tunnelWatch)

	//连接接口
	modelLink := reflect.TypeOf(model.Link{})
	app.POST("/link/list", curdApiList(modelLink))
	app.GET("/link/:id", curdApiGet(modelLink))
	app.POST("/link/:id", curdApiModify(modelLink,
		[]string{"name", "devices", "template_id",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, nil))
	app.GET("/link/:id/delete", curdApiDelete(modelLink, nil, afterLinkDelete))
	app.GET("/link/:id/stop", parseParamId, linkClose)
	app.GET("/link/:id/enable", curdApiDisable(modelLink, false, nil, nil))
	app.GET("/link/:id/disable", curdApiDisable(modelLink, true, nil, afterLinkDisable))
	app.GET("/link/:id/watch", parseParamId, linkWatch)

	//连接接口
	modelEvent := reflect.TypeOf(model.Event{})
	app.POST("/event/list", curdApiList(modelEvent))
	app.POST("/event/clear", eventClear)
	app.GET("/event/:id/delete", curdApiDelete(modelEvent, nil, nil))

	//连接接口
	modelPipe := reflect.TypeOf(model.Pipe{})
	app.POST("/pipe/list", curdApiList(modelPipe))
	app.POST("/pipe/create", curdApiCreate(modelPipe, nil, afterPipeCreate))
	app.GET("/pipe/:id", curdApiGet(modelPipe))
	app.POST("/pipe/:id", curdApiModify(modelPipe,
		[]string{"name", "devices", "template_id",
			"hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"},
		nil, afterPipeUpdate))
	app.GET("/pipe/:id/delete", curdApiDelete(modelPipe, nil, afterPipeDelete))

	app.GET("/pipe/:id/enable", curdApiDisable(modelPipe, false, nil, afterPipeEnable))
	app.GET("/pipe/:id/disable", curdApiDisable(modelPipe, true, nil, afterPipeDisable))
	app.GET("/pipe/:id/start", parseParamId, pipeStart)
	app.GET("/pipe/:id/stop", parseParamId, pipeStop)

	//系统接口
	app.GET("/system/version")
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
