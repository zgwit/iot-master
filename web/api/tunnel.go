package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3/q"
	"golang.org/x/net/websocket"
)

func tunnelRoutes(app *gin.RouterGroup) {
	app.POST("list", tunnelList)
	app.POST("create", tunnelCreate)

	app.GET("event/clear", tunnelEventClearAll)

	app.Use(parseParamId)
	app.POST(":id/update", tunnelUpdate)
	app.GET(":id/delete", tunnelDelete)
	app.GET(":id/start", tunnelStart)
	app.GET(":id/stop", tunnelStop)
	app.GET(":id/enable", tunnelEnable)
	app.GET(":id/disable", tunnelDisable)
	app.GET(":id/watch", tunnelWatch)
	app.POST(":id/event/list", tunnelEvent)
	app.GET(":id/event/clear", tunnelEventClear)
}

func tunnelList(ctx *gin.Context) {
	tunnels, cnt, err := normalSearch(ctx, database.Master, &model.Tunnel{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, tunnels, cnt)
}

func tunnelCreate(ctx *gin.Context) {
	var tunnel model.Tunnel
	err := ctx.ShouldBindJSON(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, tunnel)

	//启动
	//tunnelStart(ctx)
	go func() {
		err := master.LoadTunnel(tunnel.ID)
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func tunnelUpdate(ctx *gin.Context) {
	var tunnel model.Tunnel
	err := ctx.ShouldBindJSON(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}
	tunnel.ID = ctx.GetInt("id")

	err = database.Master.Update(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, tunnel)

	//重新启动
	go func() {
		tnl := master.GetTunnel(ctx.GetInt("id"))
		if tnl == nil {
			return
		}
		err = tnl.Instance.Close()
		if err != nil {
			log.Error(err)
			return
		}
		err = tnl.Instance.Open()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func tunnelDelete(ctx *gin.Context) {
	tunnel := model.Tunnel{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, tunnel)

	//关闭
	go func() {
		tunnel := master.GetTunnel(ctx.GetInt("id"))
		if tunnel == nil {
			return
		}
		err := tunnel.Instance.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func tunnelStart(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt("id"))
	if tunnel == nil {
		replyFail(ctx, "not found")
		return
	}
	err := tunnel.Instance.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func tunnelStop(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt("id"))
	if tunnel == nil {
		replyFail(ctx, "not found")
		return
	}
	err := tunnel.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}


func tunnelEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Tunnel{ID: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//启动
	go func() {
		tunnel := master.GetTunnel(ctx.GetInt("id"))
		if tunnel == nil {
			return
		}
		err := tunnel.Instance.Open()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func tunnelDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Tunnel{ID: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		tunnel := master.GetTunnel(ctx.GetInt("id"))
		if tunnel == nil {
			return
		}
		err := tunnel.Instance.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func tunnelWatch(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt("id"))
	if tunnel == nil {
		replyFail(ctx, "找不到通道")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, tunnel.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func tunnelEvent(ctx *gin.Context) {
	events, cnt, err := normalSearchById(ctx, database.History, "TunnelID", ctx.GetInt("id"), &model.TunnelEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, events, cnt)
}

func tunnelEventClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("TunnelID", ctx.GetInt("id"))).Delete(&model.TunnelEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func tunnelEventClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.TunnelEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
